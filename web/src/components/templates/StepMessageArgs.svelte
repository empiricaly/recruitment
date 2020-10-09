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
  import Select from "../base/Select.svelte";
  import Label from "../base/Label.svelte";
  import Input from "../base/Input.svelte";
  import SlideOver from "../overlays/SlideOver.svelte";
  import CodeMirror from "../editors/CodeMirror.svelte";
  import { uniqueID } from "../../utils/uniq.js";

  export let msgArgs;
  let showVariables = false;

  const uniq = uniqueID();
</script>

<div class="">
  <Label
    forID={uniq('subject')}
    text="Message Subject"
    question="The subject line of the email message to send" />
  <Input
    max={200}
    id={uniq('subject')}
    bind:value={msgArgs.subject}
    placeholder="Message Subject" />
</div>
<div class="mt-4">
  <Label
    forID={uniq('url')}
    text="Target URL"
    optional
    question="URL Participants should be forwarded to" />
  <Input
    id={uniq('url')}
    bind:value={msgArgs.url}
    placeholder="https://experiment.example.com" />
</div>
<div class="mt-4">
  <div class="flex justify-between">
    <Label
      forID={uniq('message')}
      text="Message Template"
      question="Text body of HIT" />

    <div class="flex text-gray-300 text-sm items-baseline">
      <button class="mr-2" on:click={() => (showVariables = true)}>
        variables
      </button>
      •
      <div class="w-32 flex-shrink-0">
        <Select
          thin
          bind:value={msgArgs.messageType}
          options={modes}
          placeholder="Mode" />
      </div>
    </div>
  </div>
  <div class="border">
    <CodeMirror
      bind:value={msgArgs.message}
      mode={typeToMode[msgArgs.messageType]} />
  </div>

  <SlideOver title="Message Template Variables" bind:open={showVariables}>
    <div class="mr-6 text-gray-400 text-sm">
      The message can be written in Markdown, HTML, React or Svelte. The
      template is given variables that can be used in the message. That availble
      variables are:
      <ul class="ml-5 list-outside list-disc">
        <li>
          <code>url:</code>
          The target URL passed above. The actual URL passed to the template
          will be a unique redirect URL for each participant. This allows
          tracking of page loads.
        </li>
        <li>
          <code>step:</code>
          The current Step object, which contains all the configuration added
          here. It also points to it's parent Template object. See documentation
          for further details.
        </li>
        <li>
          <code>stepRun:</code>
          The current Step Run object, which contains the current step's run
          information. It also points to it's parent's Run object. See
          documentation for further details.
        </li>
        <li><code>participant:</code> The current participant.</li>
      </ul>
    </div>
  </SlideOver>
</div>
