<script>
  import ParticipantList from "../components/participants/ParticipantList.svelte";
  import Import from "../components/participants/Import.svelte";
  import Layout from "../layouts/Layout.svelte";
  import { exportParticipants } from "../lib/models/participants/participants.js";
  import { GET_PROJECT_PARTICIPANTS } from "../lib/queries";
  import { notify } from "../components/overlays/Notification.svelte";

  const queryArgs = (project) => ({
    query: GET_PROJECT_PARTICIPANTS,
    variables: { projectID: project.projectID },
  });

  let participants;
  let keys;
  let projectID;
  let isOpen = false;
  let loading = false;

  function setLoading(val) {
    loading = val;
  }

  async function handleClick(event) {
    const { action, project } = event.detail;
    switch (action) {
      case "exportjson":
      case "exportcsv":
        setLoading(true);
        notify({
          failed: false,
          title: `Exporting participants.`,
        });
        exportParticipants({
          project,
          keys,
          type: action === "exportjson" ? "json" : "csv",
          setLoading,
        });
        break;
      case "import":
        isOpen = true;
        projectID = project.id;
        break;
      default:
        console.error(`Unknown action: ${action}`);
        break;
    }
  }

  let actions = [];

  $: {
    actions = [];

    actions.push({
      text: "Import",
      action: "import",
      disabled: Boolean(loading),
      icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M48.048 304h73.798v128c0 26.51 21.49 48 48 48h108.308c26.51 0 48-21.49 48-48V304h73.789c42.638 0 64.151-51.731 33.941-81.941l-175.943-176c-18.745-18.745-49.137-18.746-67.882 0l-175.952 176C-16.042 252.208 5.325 304 48.048 304zM224 80l176 176H278.154v176H169.846V256H48L224 80z"/></svg>`,
      primary: false,
    });

    actions.push({
      text: "Export CSV",
      action: "exportcsv",
      disabled: Boolean(loading || !participants || participants.length === 0),
      icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M400 208h-73.8V80c0-26.5-21.5-48-48-48H169.8c-26.5 0-48 21.5-48 48v128H48.1c-42.6 0-64.2 51.7-33.9 81.9l175.9 176c18.7 18.7 49.1 18.7 67.9 0l176-176c30-30.1 8.7-81.9-34-81.9zM224 432L48 256h121.8V80h108.3v176H400L224 432z"/></svg>`,
      primary: true,
    });

    actions.push({
      text: "Export JSON",
      action: "exportjson",
      disabled: Boolean(loading || !participants || participants.length === 0),
      icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M400 208h-73.8V80c0-26.5-21.5-48-48-48H169.8c-26.5 0-48 21.5-48 48v128H48.1c-42.6 0-64.2 51.7-33.9 81.9l175.9 176c18.7 18.7 49.1 18.7 67.9 0l176-176c30-30.1 8.7-81.9-34-81.9zM224 432L48 256h121.8V80h108.3v176H400L224 432z"/></svg>`,
      primary: true,
    });
  }
</script>

<Layout
  title="Project Participants"
  let:project
  on:click={handleClick}
  {actions}>
  <ParticipantList
    type="project"
    queryArgs={queryArgs(project)}
    bind:participants
    bind:keys />
  <Import {projectID} bind:isOpen />
</Layout>
