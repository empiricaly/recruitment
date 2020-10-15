<script>
  import TemplateSection from "../templates/TemplateSection.svelte";
  import StepRun from "./StepRun.svelte";

  export let project;
  export let run;

  console.log(project);
</script>

<TemplateSection title="Participant Selection" description="" current>
  <div class="">
    {#if run.template.selectionType === 'INTERNAL_DB'}
      {#if run.template.internalCriteria.all}
        Participants are randomly selected from the internal database.
      {:else}
        A custom subselection of Participants are randomly selected from the
        internal database.
      {/if}
    {:else if run.template.selectionType === 'MTURK_QUALIFICATIONS'}
      MTurk selection...
    {:else}Unknow Particpant Selection Type{/if}
    <div>
      <div class="">{run.template.participantCount} Participants requested</div>
    </div>
  </div>
</TemplateSection>

{#each run.template.steps as step}
  <StepRun
    {step}
    stepLength={run.template.steps.length}
    error={run.template.selectionType === 'MTURK_QUALIFICATIONS' && step.index === 0 && step.type !== 'MTURK_HIT' ? 'First step of a Run using MTurk Qualifications must be an MTurk Hit.' : ''} />
{/each}
