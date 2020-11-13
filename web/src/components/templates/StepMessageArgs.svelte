<script context="module">
  const modes = [
    {
      label: "Plain Text",
      value: "PLAIN",
    },
  ];
  const richModes = [
    ...modes,
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
    PLAIN: "plain",
    MARKDOWN: "markdown",
    HTML: "html",
    REACT: "jsx",
    SVELTE: "svelte",
  };
</script>

<script>
  import { uniqueID } from "../../utils/uniq.js";
  import Input from "../base/Input.svelte";
  import Label from "../base/Label.svelte";
  import Select from "../base/Select.svelte";
  import CodeMirror from "../editors/CodeMirror.svelte";
  import SlideOver from "../overlays/SlideOver.svelte";

  export let msgArgs;
  export let hasSubject = false;
  export let rich = false;
  let showVariables = false;

  const uniq = uniqueID();

  const modesSelect = rich ? richModes : modes;
</script>

{#if hasSubject}
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
{/if}
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
          options={modesSelect}
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
      <p class="mt-2">
        Messages can be written in Plain Text, and if it's a HIT it can also be
        Markdown or HTML.
      </p>
      <p class="mt-2">
        The template is given variables that can be used in the message. That
        availble variables are:
      </p>
      <dl class="mt-3 ml-5 list-outside list-disc">
        <dt class="mt-2"><code class="text-gray-700">url</code></dt>
        <dd class="mt-1">
          The target URL passed above. The actual URL passed to the template
          will be a unique redirect URL for each participant.
        </dd>
        <dt class="mt-2"><code class="text-gray-700">currentStep</code></dt>
        <dd class="mt-1">
          The current Step object, which contains all the configuration added
          here.
        </dd>
        <!-- <dt class="mt-2"><code class="text-gray-700">stepRun</code></dt>
        <dd class="mt-1">
          The current Step Run object, which contains the current step's run
          information. It also points to it's parent's Run object. See
          documentation for further details.
        </dd> -->
        <dt class="mt-2"><code class="text-gray-700">participant</code></dt>
        <dd class="mt-1">The current participant.</dd>
      </dl>
    </div>
  </SlideOver>
</div>
