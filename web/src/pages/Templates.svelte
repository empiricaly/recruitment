<script>
  import { mutate } from "svelte-apollo";
  import { notify } from "../components/overlays/Notification.svelte";
  import TemplateLine from "../components/templates/TemplateLine.svelte";
  import Layout from "../layouts/Layout.svelte";
  import { client } from "../lib/apollo";
  import { CREATE_TEMPLATE } from "../lib/queries";
  import { handleErrorMessage } from "../utils/errorQuery";

  export let params;

  async function handleCreate() {
    try {
      await mutate(client, {
        mutation: CREATE_TEMPLATE,
        variables: {
          input: {
            projectID: params.projectID,
            name: "",
            selectionType: "INTERNAL_DB",
            participantCount: 100,
            adult: false,
            sandbox: false,
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
        },
      });
    } catch (error) {
      handleErrorMessage(error);
      notify({
        failed: true,
        title: `Could not create Template`,
        body:
          "Something happened on the server, and we could not create a new Template as requested.",
      });
    }
  }
</script>

<Layout
  title="Templates"
  action="New Template"
  on:click={handleCreate}
  let:project>
  <div class="bg-white shadow sm:rounded-md">
    <ul>
      <TemplateLine
        projectID={project.projectID}
        templateID="2345678"
        name="Speed Dating Pilot 2"
        type="internal"
        stepCount={2}
        runCount={0}
        creator="Nicolas" />
      <TemplateLine
        projectID={project.projectID}
        templateID="2345678"
        index="1"
        name="Speed Dating Pilot 2"
        type="internal"
        stepCount={2}
        runCount={0}
        creator="Nicolas" />
      <TemplateLine
        projectID={project.projectID}
        templateID="2345678"
        index="2"
        name="Speed Dating Pilot 1"
        type="mturk"
        stepCount={4}
        runCount={5}
        creator="Nicolas" />
    </ul>
  </div>
</Layout>
