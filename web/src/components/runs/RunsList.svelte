<script>
  import { query } from "svelte-apollo";
  import { client } from "../../lib/apollo";
  import { GET_RUNS } from "../../lib/queries";
  import Callout from "../base/Callout.svelte";
  import RunLine from "./RunLine.svelte";

  export let project;
  export let limit = 0;

  $: runs = query(client, {
    query: GET_RUNS,
    variables: { projectID: project.projectID, limit },
  });
</script>

{#await $runs}
  Loading...
{:then result}
  <div>
    {#if result.data.project.runs.length === 0}
      <Callout color="yellow">
        You have no runs yet.
        <br />
        Create one now!
      </Callout>
    {/if}
  </div>

  <div class="bg-white shadow sm:rounded-md">
    <ul>
      {#each result.data.project.runs as run, index (run.id)}
        <RunLine
          projectID={project.id}
          projectName={project.projectID}
          runID={run.id}
          startAt={run.startAt}
          startedAt={run.startedAt}
          endedAt={run.endedAt}
          name={run.name}
          {index}
          stepCount={run.template.steps.length}
          status={run.status} />
      {/each}
    </ul>
  </div>
{:catch error}
  Error loading Runs:
  {error}
{/await}
