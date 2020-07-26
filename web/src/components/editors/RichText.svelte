<script>
  import EditorJS from "@editorjs/editorjs";
  import Header from "@editorjs/header";
  import List from "@editorjs/list";

  export let data = {};
  export let placeholder = "";

  function init(el) {
    const editor = new EditorJS({
      holder: el,
      tools: {
        header: Header,
        list: List
      },
      placeholder,
      data
    });
  }
</script>

<div class="editor" />
