<script context="module">
  const modes = [
    {
      label: "Markdown",
      value: "MARKDOWN",
    },
    {
      label: "HTML",
      value: "HTML",
    },
    {
      label: "React",
      value: "REACT",
    },
    {
      label: "Svelte",
      value: "SVELTE",
    },
  ];

  const typeToMode = {
    MARKDOWN: "markdown",
    HTML: "html",
    REACT: "jsx",
    SVELTE: "svelte",
  };
</script>

<script>
  import { uniqueID } from "../../utils/uniq.js";
  import Label from "../base/Label.svelte";
  import CodeMirror from "../editors/CodeMirror.svelte";

  export let msgArgs;
  export let hasSubject = false;
  let showVariables = false;

  const uniq = uniqueID();
</script>

{#if hasSubject}
  <div class="">
    <Label
      forID={uniq('subject')}
      text="Message Subject"
      question="The subject line of the email message to send" />
    {msgArgs.subject}
  </div>
{/if}

{#if msgArgs.url}
  <div class="mt-4">
    <Label
      forID={uniq('url')}
      text="Target URL"
      question="URL Participants should be forwarded to" />
    {msgArgs.url}
  </div>
{/if}

<div class="mt-4">
  <div class="flex justify-between">
    <Label
      forID={uniq('message')}
      text="Message Template"
      question="Text body of HIT" />

    <div class="flex text-gray-300 text-sm items-baseline">
      {msgArgs.messageType}
    </div>
  </div>
  <div class="mt-2 border">
    <CodeMirror
      readonly
      value={msgArgs.message}
      mode={typeToMode[msgArgs.messageType]} />

    <!-- {msgArgs.message} -->
  </div>
</div>
