<script>
  import dayjs from "dayjs";
  import { timer } from "../../utils/timer.js";
  import StepRunMTurkHit from "./StepRunMTurkHit.svelte";
  import StepRunMTurkMessage from "./StepRunMTurkMessage.svelte";
  import StepRunParticipantFilter from "./StepRunParticipantFilter.svelte";
  import StepRunTemplateSection from "./StepRunTemplateSection.svelte";

  export let step;
  export let stepRun;
  $: current = stepRun && stepRun.status === "RUNNING";

  let finishedAt;
  let remaining;
  let remainingPercent = 0;
  let localTimer;

  $: if (stepRun && stepRun.startedAt && !stepRun.endedAt) {
    localTimer = $timer;
  }

  $: if (stepRun && stepRun.endedAt) {
    remaining = null;
    remainingPercent = 0;
    finishedAt = dayjs(stepRun.endedAt).calendar();
  }

  $: if (localTimer) {
    remaining = remaining = dayjs.duration(
      dayjs(stepRun.startedAt).add(step.duration, "minutes").diff(dayjs())
    );
    remainingPercent = (100 / (step.duration * 60)) * remaining.as("seconds");
  }

  let showDetails = false;
  function handleShowDetails() {
    console.log("SHOW DETAILS", showDetails);
    showDetails = !showDetails;
  }
</script>

<StepRunTemplateSection header {current} progress={remainingPercent}>
  <div slot="title" class="flex">
    <div class="font-semibold {current ? 'text-mint-800' : 'text-gray-400'}">
      Step
      {step.index + 1}
    </div>
    <div class="mx-2 {current ? 'text-mint-300' : 'text-gray-300'}">
      <!-- ·• -->
      →
    </div>
    <div class={current ? 'text-mint-700' : 'text-gray-400'}>
      {#if step.type === 'MTURK_HIT'}
        MTurk HIT
      {:else if step.type === 'MTURK_MESSAGE'}
        MTurk Message
      {:else if step.type === 'PARTICIPANT_FILTER'}Participant Filter{/if}
    </div>
  </div>
  <div slot="description">
    <!-- Add description -->
  </div>

  <div slot="header" class="cursor-pointer" on:click={handleShowDetails}>
    <div class="flex justify-between">
      <div class="mr-2 flex items-baseline">
        {#if finishedAt}
          <div class="text-gray-500  mr-2">Finished</div>
          {finishedAt}
        {:else if remaining}
          <div class="{current ? 'text-mint-300' : 'text-gray-500'}  mr-2">
            Remaining
          </div>
          {remaining.humanize()}
        {/if}
      </div>

      <div class="ml-2 flex items-center md:justify-end">
        <div class="{current ? 'text-mint-300' : 'text-gray-500'} mr-2">
          Duration
        </div>
        {step.duration}
        minutes
      </div>
    </div>
  </div>

  {#if showDetails}
    <div class="px-4 py-5 sm:p-6">
      {#if step.type === 'MTURK_HIT'}
        <StepRunMTurkHit bind:step />
      {:else if step.type === 'MTURK_MESSAGE'}
        <StepRunMTurkMessage bind:step />
      {:else if step.type === 'PARTICIPANT_FILTER'}
        <StepRunParticipantFilter bind:step />
      {:else}Unknown Step Type?!?{/if}
    </div>
  {/if}
</StepRunTemplateSection>
