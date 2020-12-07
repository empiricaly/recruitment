<script>
  import Label from "../base/Label.svelte";
  import StepRunMessageArgs from "./StepRunMessageArgs.svelte";

  export let step;
</script>

<div class="md:grid grid-cols-3 gap-6">
  <div class="col-span-2">
    <div class="">
      <Label
        text="HIT Title"
        question={`Describe the task to Workers. Be as specific as possible,
  e.g. "answer a survey about movies", instead of "short survey", so Workers
  know what to expect.
  Tasks that contain adult content are required to include the following phrase
  in your task title: (WARNING: This HIT may contain adult content. Worker
  discretion is advised.)`} />
      {step.hitArgs.title}
    </div>

    <div class="mt-4">
      <Label
        text="HIT Description"
        question="Give more detail about this task. This gives Workers a bit
        more information before they decide to view your task" />
      {step.hitArgs.description}
    </div>

    <div class="mt-4">
      <Label
        text="HIT Keywords"
        question="Provide keywords that will help Workers search for your tasks" />
      {step.hitArgs.keywords}
    </div>
  </div>

  <div class="mt-4 md:mt-0">
    <div class="">
      <Label text="HIT Reward" question="MTurk HIT reward for task in USD" />
      ${step.hitArgs.reward}
      USD
    </div>

    <div class="mt-4">
      <Label
        text="HIT Accepted HIT Timeout"
        question="Timeout of a single accepted HIT in seconds" />
      {step.hitArgs.timeout}
      minutes
    </div>

    <!-- <div class="mt-4">
      <Label
        text="Number of HITs to publish"
        question="Maximum number of HITs to publish." />
      {step.hitArgs.workersCount}
      workers
    </div> -->
  </div>
</div>

<div class="mt-2 hidden sm:block">
  <div class="py-5">
    <div class="border-t border-gray-200" />
  </div>
</div>

<div class="mt-1">
  <StepRunMessageArgs msgArgs={step.msgArgs} />
</div>
