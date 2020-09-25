<script>
  import { query } from "svelte-apollo";
  import { client } from "../../lib/apollo";
  import { GET_RUNS } from "../../lib/queries";
  import RunLine from "./RunLine.svelte";

  export let project;

  $: runs = query(client, {
    query: GET_RUNS,
    variables: { projectID: project.projectID },
  });
</script>

{#await $runs}
  Loading...
{:then result}
  <div>
    <div class="bg-white shadow sm:rounded-md">
      <ul>
        {#each result.data.project.runs as run, index (run.id)}
          <RunLine
            projectID={project.projectID}
            runID={run.id}
            startAt={run.startAt}
            startedAt={run.startedAt}
            endedAt={run.endedAt}
            name={run.name}
            {index}
            stepCount={run.procedure.steps.length}
            status={run.status} />
        {/each}
      </ul>
    </div>
  </div>
{:catch error}
  Error loading Runs: {error}
{/await}
