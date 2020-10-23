<script>
  import ParticipantList from "../components/participants/ParticipantList.svelte";
  import Layout from "../layouts/Layout.svelte";
  import { GET_PROJECT_PARTICIPANTS } from "../lib/queries";
</script>

<Layout title="Project Participants" let:project>
  <ParticipantList
    type="project"
    queryArgs={{ query: GET_PROJECT_PARTICIPANTS, variables: { projectID: project.projectID } }} />
</Layout>
