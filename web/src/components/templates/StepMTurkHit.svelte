<script context="module">
  const modes = [
    {
      label: "Markdown",
      value: "markdown"
    },
    {
      label: "HTML",
      value: "html"
    },
    {
      label: "React",
      value: "jsx"
    },
    {
      label: "Svelte",
      value: "svelte"
    }
  ];
</script>

<script>
  import Select from "../base/Select.svelte";
  import Label from "../base/Label.svelte";
  import Input from "../base/Input.svelte";
  import Textarea from "../base/Textarea.svelte";
  import SlideOver from "../overlays/SlideOver.svelte";
  import CodeMirror from "../editors/CodeMirror.svelte";
  import Button from "../base/Button.svelte";
  import StepMessageArgs from "./StepMessageArgs.svelte";
  import { uniqueID } from "../../utils/uniq.js";

  export let step;
  let mode = "markdown";
  let showVariables = false;

  const uniq = uniqueID();

  $: hitArgs = step.args[0];
  $: msgArgs = step.args[1];
</script>

<div class="md:grid grid-cols-3 gap-6">
  <div class="col-span-2">
    <div class="">
      <Label
        forID={uniq('title')}
        text="HIT Title"
        question={`Describe the task to Workers. Be as specific as possible,
  e.g. "answer a survey about movies", instead of "short survey", so Workers
  know what to expect.
  Tasks that contain adult content are required to include the following phrase
  in your task title: (WARNING: This HIT may contain adult content. Worker
  discretion is advised.)`} />
      <Input
        id={uniq('title')}
        bind:value={hitArgs.title}
        required
        placeholder="Describe the task to Workers" />
    </div>

    <div class="mt-4 ">
      <Label
        forID={uniq('description')}
        text="HIT Description"
        question="Give more detail about this task. This gives Workers a bit
        more information before they decide to view your task" />
      <Textarea
        id={uniq('description')}
        bind:value={hitArgs.description}
        required
        placeholder="Give more detail about this task" />
    </div>

    <div class="mt-4 ">
      <Label
        forID={uniq('keywords')}
        text="HIT Keywords"
        question="Provide keywords that will help Workers search for your tasks" />
      <Input
        id={uniq('keywords')}
        bind:value={hitArgs.keywords}
        required
        placeholder="Comma-Seperated Keywords" />
    </div>
  </div>

  <div class="mt-4 md:mt-0">
    <div class="">
      <Label
        forID={uniq('reward')}
        text="HIT Reward"
        question="MTurk HIT reward for task in USD" />
      <Input
        id={uniq('reward')}
        type="number"
        min="0"
        left="$"
        right="USD"
        bind:value={hitArgs.reward}
        inputmode="numeric"
        required
        placeholder="0.0" />
    </div>

    <div class="mt-4">
      <Label
        forID={uniq('timeout')}
        text="HIT Accepted HIT Timeout"
        question="Timeout of a single accepted HIT in seconds" />
      <Input
        type="number"
        id={uniq('timeout')}
        right="seconds"
        bind:value={hitArgs.timeout}
        inputmode="numeric"
        required
        placeholder="0" />
    </div>
  </div>
</div>

<div class="mt-2 hidden sm:block">
  <div class="py-5">
    <div class="border-t border-gray-200" />
  </div>
</div>

<div class="mt-1">
  <StepMessageArgs bind:msgArgs />
</div>
