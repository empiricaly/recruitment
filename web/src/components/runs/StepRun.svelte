<script context="module">
  const stepTypeName = {
    MTURK_HIT: "MTurk HIT Step",
    MTURK_MESSAGE: "MTurk Message Step",
    PARTICIPANT_FILTER: "Participant Filter Step",
  };
  const stepTypeDesc = {
    MTURK_HIT: "TODO: Add Description",
    MTURK_MESSAGE: "TODO: Add Description",
    PARTICIPANT_FILTER: "TODO: Add Description",
  };
</script>

<script>
  import TemplateSection from "../templates/TemplateSection.svelte";
  import StepRunMTurkHit from "./StepRunMTurkHit.svelte";
  import StepRunMTurkMessage from "./StepRunMTurkMessage.svelte";
  import StepRunParticipantFilter from "./StepRunParticipantFilter.svelte";

  export let step;
  // export let stepLength;
</script>

<TemplateSection title={stepTypeName[step.type]} header>
  <div slot="header">
    <div class="flex justify-between">
      <div class="mr-2 flex items-baseline">
        <div class="text-mint-300 mr-2">Total Duration</div>
        {step.duration}
        minutes
      </div>

      <div class="ml-2 flex items-center md:justify-end">
        <div class="text-mint-300 mr-2">Remaining</div>
        {step.duration}
        minutes
      </div>
    </div>
  </div>
  <div slot="description">{stepTypeDesc[step.type]}</div>
  {#if step.type === 'MTURK_HIT'}
    <StepRunMTurkHit bind:step />
  {:else if step.type === 'MTURK_MESSAGE'}
    <StepRunMTurkMessage bind:step />
  {:else if step.type === 'PARTICIPANT_FILTER'}
    <StepRunParticipantFilter bind:step />
  {:else}Unknown Step Type?!?{/if}
</TemplateSection>
