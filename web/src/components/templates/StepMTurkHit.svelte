<script>
  import { uniqueID } from "../../utils/uniq.js";
  import Input from "../base/Input.svelte";
  import Label from "../base/Label.svelte";
  import Textarea from "../base/Textarea.svelte";
  import StepMessageArgs from "./StepMessageArgs.svelte";

  export let step;
  let mode = "plain";
  let showVariables = false;

  const uniq = uniqueID();
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
        bind:value={step.hitArgs.title}
        required
        placeholder="Describe the task to Workers" />
    </div>

    <div class="mt-4">
      <Label
        forID={uniq('description')}
        text="HIT Description"
        question="Give more detail about this task. This gives Workers a bit
        more information before they decide to view your task" />
      <Textarea
        id={uniq('description')}
        bind:value={step.hitArgs.description}
        required
        placeholder="Give more detail about this task" />
    </div>

    <div class="mt-4">
      <Label
        forID={uniq('keywords')}
        text="HIT Keywords"
        question="Provide keywords that will help Workers search for your tasks" />
      <Input
        id={uniq('keywords')}
        bind:value={step.hitArgs.keywords}
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
        bind:value={step.hitArgs.reward}
        inputmode="numeric"
        required
        placeholder="0.0" />
    </div>

    <div class="mt-4">
      <Label
        forID={uniq('timeout')}
        text="HIT Accepted HIT Timeout"
        question="Timeout of a single accepted HIT in minutes" />
      <Input
        type="number"
        id={uniq('timeout')}
        right="minutes"
        bind:value={step.hitArgs.timeout}
        inputmode="numeric"
        required
        placeholder="0" />
    </div>

    <!-- <div class="mt-4">
      <Label
        forID={uniq('workersCount')}
        text="Number of HITs to publish"
        question="Maximum number of HITs to publish." />
      <Input
        type="number"
        id={uniq('workersCount')}
        right="hits"
        bind:value={step.hitArgs.workersCount}
        inputmode="numeric"
        required
        placeholder="0" />
    </div> -->
  </div>
</div>

<div class="mt-2 hidden sm:block">
  <div class="py-5">
    <div class="border-t border-gray-200" />
  </div>
</div>

<div class="mt-1">
  <StepMessageArgs bind:msgArgs={step.msgArgs} rich />
</div>
