<script>
  import dayjs from "dayjs";
  import { createEventDispatcher } from "svelte";
  import { query } from "svelte-apollo";
  import { client } from "../../lib/apollo";
  import { GET_RUNNING_RUN, GET_RUN_PARTICIPANTS } from "../../lib/queries";
  import { handleErrorMessage } from "../../utils/errorQuery";
  import { timer } from "../../utils/timer.js";
  import ParticipantList from "../participants/ParticipantList.svelte";
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

  $: try {
    $runningRun.then((result) => {
      runB = result.data.project.runs[0];
    });
  } catch (error) {
    handleErrorMessage(error);
  }

  $: if (runB && runB.status == "RUNNING") {
    if ($timer) {
      if ($runningRun) {
        runningRun.refetch();
      }
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
  <StepRun
    {run}
    steps={run.template.steps}
    {step}
    stepRun={runB && runB.steps[index]} />
{/each}

{#if run && run.status === 'DONE'}
  <div class="mt-8 mb-6 hidden sm:block">
    <div class="py-5">
      <div class="border-t border-gray-200" />
    </div>
  </div>

  <ParticipantList
    queryArgs={{ query: GET_RUN_PARTICIPANTS, variables: { projectID: project.projectID, runID: run.id } }} />
{:else if run && run.status === 'FAILED'}
  <div class="mt-8 hidden sm:block">
    <div class="py-5">
      <div class="border-t border-gray-200" />
    </div>
  </div>

  <div class="text-red-700 mt-8 p-4 border border-red-600 bg-red-100 rounded">
    <details class="">
      <summary class="font-medium outline-none">
        Run encountered an error and could not continue
      </summary>

      <p class="mt-2 italic">
        If this does not make sense, please create an issue in Github
      </p>

      <div
        class="mt-2 p-3 bg-gray-100 shadow-inner rounded text-gray-800 border">
        <code>
          <pre class="overflow-x-auto">
            {run.error}
          </pre>
        </code>
      </div>
    </details>
  </div>
{/if}
