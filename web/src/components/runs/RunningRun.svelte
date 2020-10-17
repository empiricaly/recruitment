<script>
  import dayjs from "dayjs";
  import { createEventDispatcher } from "svelte";
  import { query } from "svelte-apollo";
  import { client } from "../../lib/apollo";
  import { GET_RUNNING_RUN } from "../../lib/queries";
  import { timer } from "../../utils/timer.js";
  import TemplateSection from "../templates/TemplateSection.svelte";
  import StepRun from "./StepRun.svelte";

  const dispatch = createEventDispatcher();

  export let project;
  export let run;

  let runB = null;

  console.log(project, run);

  const runningRun = query(client, {
    query: GET_RUNNING_RUN,
    variables: { projectID: project.projectID, runID: run.id },
  });

  $: $runningRun.then((result) => {
    runB = result.data.project.runs[0];
  });

  $: if ($timer) {
    if ($runningRun) {
      runningRun.refetch();
    }
  }

  $: if (runB && runB.status == "DONE") {
    dispatch("refresh");
  }
</script>

{#if run.status === 'CREATED' && run.startAt}
  <TemplateSection title="" description="" current>
    <div class="">Scheduled for {dayjs(run.startAt).calendar()}</div>
  </TemplateSection>
{/if}

<TemplateSection title="Participant Selection" description="">
  <div class="">
    {#if run.template.selectionType === 'INTERNAL_DB'}
      {#if run.template.internalCriteria.all}
        Participants are randomly selected from the internal database.
      {:else}
        A custom selection of Participants are randomly selected from the
        internal database.
      {/if}
    {:else if run.template.selectionType === 'MTURK_QUALIFICATIONS'}
      Participants are coming from MTurk.
    {:else}Unknow Particpant Selection Type{/if}
    {run.template.participantCount}
    Participants requested.
  </div>
</TemplateSection>

{#each run.template.steps as step, index}
  <StepRun {step} stepRun={runB && runB.steps[index]} />
{/each}
