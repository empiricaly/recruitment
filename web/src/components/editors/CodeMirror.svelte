<script context="module">
  import "codemirror/mode/javascript/javascript.js";
  import "codemirror/mode/htmlmixed/htmlmixed.js";
  import "codemirror/mode/markdown/markdown.js";
  import "codemirror/mode/handlebars/handlebars.js";
  import "codemirror/mode/jsx/jsx.js";
  import "codemirror/addon/edit/closebrackets.js";
  import "codemirror/addon/edit/closetag.js";
  import "codemirror/addon/comment/comment.js";
</script>

<script>
  import { onMount, createEventDispatcher } from "svelte";
  import CodeMirror from "codemirror";

  const dispatch = createEventDispatcher();

  export let value = "";
  export let mode = "javascript";
  export let readonly = false;
  export let errorLoc = null;
  export let lineNumbers = true;

  let w;
  let h;

  let prevMode = mode;
  $: {
    if (prevMode !== mode) {
      console.log("updating mode");
      createEditor(mode);
      prevMode = mode;
    }
  }

  export function update(new_value) {
    value = new_value;
    if (editor) {
      const { left, top } = editor.getScrollInfo();
      editor.setValue((value = new_value));
      editor.scrollTo(left, top);
    }
  }

  export function resize() {
    editor.refresh();
  }

  export function focus() {
    editor.focus();
  }

  const modes = {
    markdown: {
      name: "markdown",
      base: "text/x-markdown"
    },
    jsx: {
      name: "jsx"
    },
    html: {
      name: "htmlmixed",
      base: "text/html"
    },
    svelte: {
      name: "handlebars",
      base: "text/html"
    }
  };

  const refs = {};

  let editor;
  let updating_externally = false;
  let marker;
  let error_line;
  let destroyed = false;

  $: if (editor && w && h) {
    editor.refresh();
  }

  $: {
    if (marker) marker.clear();
    if (errorLoc) {
      const line = errorLoc.line - 1;
      const ch = errorLoc.column;
      marker = editor.markText(
        { line, ch },
        { line, ch: ch + 1 },
        {
          className: "error-loc"
        }
      );
      error_line = line;
    } else {
      error_line = null;
    }
  }

  let previous_error_line;
  $: if (editor) {
    if (previous_error_line != null) {
      editor.removeLineClass(previous_error_line, "wrap", "error-line");
    }
    if (error_line && error_line !== previous_error_line) {
      editor.addLineClass(error_line, "wrap", "error-line");
      previous_error_line = error_line;
    }
  }

  onMount(() => {
    createEditor(mode || "svelte").then(() => {
      if (editor) editor.setValue(value || "");
    });

    return () => {
      destroyed = true;
      if (editor) editor.toTextArea();
    };
  });

  let first = true;
  async function createEditor(mode) {
    if (destroyed || !CodeMirror) return;
    if (editor) editor.toTextArea();

    const opts = {
      lineNumbers,
      lineWrapping: true,
      indentWithTabs: true,
      indentUnit: 2,
      tabSize: 2,
      value: "",
      mode: modes[mode] || {
        name: mode
      },
      readOnly: readonly,
      autoCloseBrackets: true,
      autoCloseTags: true,
      viewportMargin: Infinity,
      extraKeys: {
        Tab: function(cm) {
          var spaces = Array(cm.getOption("indentUnit") + 1).join(" ");
          cm.replaceSelection(spaces);
        },
        "Cmd-/": "toggleComment",
        "Ctrl-/": "toggleComment"
      }
    };

    // Creating a text editor is a lot of work, so we yield
    // the main thread for a moment. This helps reduce jank
    if (first) await sleep(50);
    if (destroyed) return;
    editor = CodeMirror.fromTextArea(refs.editor, opts);
    editor.on("change", instance => {
      if (!updating_externally) {
        value = instance.getValue();
        dispatch("change", { value });
      }
    });
    if (first) await sleep(50);
    editor.refresh();
    first = false;
  }

  function sleep(ms) {
    return new Promise(fulfil => setTimeout(fulfil, ms));
  }
</script>

<textarea tabindex="0" bind:this={refs.editor} readonly {value} />

<style>
  textarea {
    visibility: hidden;
  }
  :global(.CodeMirror) {
    height: auto;
  }
</style>
