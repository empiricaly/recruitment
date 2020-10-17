<script>
  import { query } from "svelte-apollo";
  import Run from "../components/runs/Run.svelte";
  import { client } from "../lib/apollo";
  import { GET_RUN } from "../lib/queries";
  import { params } from "../lib/routing";

  $: runID = $params.runID;
  $: projectID = $params.projectID;

  $: runs =
    runID &&
    projectID &&
    query(client, { query: GET_RUN, variables: { projectID, runID } });

  function refresh() {
    if (runs) {
      runs.refetch();
    }
  }
</script>

{#if runs}
  {#await $runs}
    Loading...
  {:then result}
    {#if result.data.project.runs.length === 0}
      Run not found!
    {:else}
      <Run
        on:refresh={refresh}
        project={result.data.project}
        run={result.data.project.runs[0]} />
    {/if}
  {:catch error}
    Error loading Run:
    {error}
  {/await}
{/if}
