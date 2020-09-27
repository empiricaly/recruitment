<script>
  import { mutate, query } from "svelte-apollo";
  import { notify } from "../components/overlays/Notification.svelte";
  import RunsList from "../components/runs/RunsList.svelte";
  import Layout from "../layouts/Layout.svelte";
  import { client } from "../lib/apollo";
  import { CREATE_RUN, GET_RUNS } from "../lib/queries";
  import { push } from "../lib/routing";

  $: runs = query(client, { query: GET_RUNS });

  $: console.log(runs);

  async function handleCreate(event) {
    const { project } = event.detail;
    console.log("project", project);
    let notifClose;
    try {
      const input = {
        projectID: project.id,
        procedure: {
          name: "New Run",
          selectionType: "INTERNAL_DB",
          participantCount: 100,
          adult: false,
          mturkCriteria: { qualifications: [] },
          internalCriteria: {
            condition: {
              and: [
                {
                  key: "",
                  comparator: "EQUAL_TO",
                  values: [],
                },
              ],
            },
          },
          steps: [],
        },
      };

      console.log(JSON.stringify(input, null, "  "));

      notifClose = notify({
        title: `Creating new Run`,
      });

      const result = await mutate(client, {
        mutation: CREATE_RUN,
        variables: {
          input,
        },
      });
      notifClose();
      notify({
        success: true,
        title: `Created new Run successfully`,
        // body:
        //   "Something happened on the server, and we could not create a new Run as requested.",
      });

      push(`/projects/${project.projectID}/runs/${result.data.createRun.id}`);
    } catch (error) {
      console.error(error);
      notifClose();
      notify({
        failed: true,
        title: `Could not create Template`,
        body:
          "Something happened on the server, and we could not create a new Run as requested.",
      });
    }
  }
</script>

<Layout title="Runs" action="New Run" on:click={handleCreate} let:project>
  <RunsList {project} />
</Layout>
