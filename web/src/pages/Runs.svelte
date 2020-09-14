<script>
  import Layout from "../layouts/Layout.svelte";
  import RunLine from "../components/runs/RunLine.svelte";
  import dayjs from "dayjs";

  export let params;
</script>

<Layout title="Runs">

  <div>
    <div class="bg-white shadow sm:rounded-md">
      <ul>
        <RunLine
          projectID={params.projectID}
          runID="2345678"
          startAt={dayjs().add(26, 'hour')}
          name="Speed Dating Pilot 2"
          index={0}
          stepCount={2}
          status="created" />
        <RunLine
          projectID={params.projectID}
          runID="2345678"
          startedAt={dayjs().subtract(10, 'minute')}
          name="Speed Dating Pilot 2"
          index={1}
          stepCount={2}
          status="running" />
        <RunLine
          projectID={params.projectID}
          runID="2345678"
          startedAt={dayjs().subtract(122, 'minute')}
          endedAt={dayjs().subtract(22, 'minute')}
          name="Speed Dating Pilot 2"
          index={1}
          stepCount={2}
          status="done" />
        <RunLine
          projectID={params.projectID}
          runID="2345678"
          endedAt={dayjs().subtract(354, 'minute')}
          name="Speed Dating Pilot 1"
          index={1}
          stepCount={4}
          status="terminated" />
        <RunLine
          projectID={params.projectID}
          runID="2345678"
          startedAt={dayjs().subtract(574, 'minute')}
          endedAt={dayjs().subtract(570, 'minute')}
          name="Speed Dating Pilot 1"
          index={1}
          stepCount={4}
          status="failed" />
      </ul>
    </div>
  </div>
</Layout>
