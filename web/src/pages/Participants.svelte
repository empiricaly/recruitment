<script>
  import dayjs from "dayjs";
  import ParticipantList from "../components/participants/ParticipantList.svelte";
  import Layout from "../layouts/Layout.svelte";
  import { participantsExportFormat } from "../lib/models/participants/participants.js";
  import { GET_PROJECT_PARTICIPANTS } from "../lib/queries";
  import { toCSV } from "../utils/csv.js";
  import { download } from "../utils/download.js";

  const queryArgs = (project) => ({
    query: GET_PROJECT_PARTICIPANTS,
    variables: { projectID: project.projectID },
  });

  let participants;
  let keys;

  function handleClick(event) {
    const { action, project } = event.detail;
    switch (action) {
      case "exportjson": {
        const out = participantsExportFormat(participants, keys);
        const content = JSON.stringify(out);
        const mime = "application/json;charset=utf-8";
        const date = dayjs().format("YYYY-MM-DDTHH:mm:ss");
        const filename = `Empirica recruitment export – ${date}.json`;
        download(content, filename, mime);
        break;
      }
      case "exportcsv": {
        const out = participantsExportFormat(participants, keys, true);
        const content = toCSV(out);
        const mime = "text/csv;charset=utf-8";
        const date = dayjs().format("YYYY-MM-DDTHH:mm:ss");
        const filename = `Empirica recruitment export – ${date}.csv`;
        download(content, filename, mime);
        break;
      }
      case "import":
        console.log("should import", project);
        break;
      default:
        console.error(`Unknown action: ${action}`);
        break;
    }
  }

  const actions = [];

  actions.push({
    text: "Import",
    action: "import",
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M48.048 304h73.798v128c0 26.51 21.49 48 48 48h108.308c26.51 0 48-21.49 48-48V304h73.789c42.638 0 64.151-51.731 33.941-81.941l-175.943-176c-18.745-18.745-49.137-18.746-67.882 0l-175.952 176C-16.042 252.208 5.325 304 48.048 304zM224 80l176 176H278.154v176H169.846V256H48L224 80z"/></svg>`,
    primary: false,
  });

  actions.push({
    text: "Export CSV",
    action: "exportcsv",
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M400 208h-73.8V80c0-26.5-21.5-48-48-48H169.8c-26.5 0-48 21.5-48 48v128H48.1c-42.6 0-64.2 51.7-33.9 81.9l175.9 176c18.7 18.7 49.1 18.7 67.9 0l176-176c30-30.1 8.7-81.9-34-81.9zM224 432L48 256h121.8V80h108.3v176H400L224 432z"/></svg>`,
    primary: true,
  });

  actions.push({
    text: "ExportJSON",
    action: "exportjson",
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M400 208h-73.8V80c0-26.5-21.5-48-48-48H169.8c-26.5 0-48 21.5-48 48v128H48.1c-42.6 0-64.2 51.7-33.9 81.9l175.9 176c18.7 18.7 49.1 18.7 67.9 0l176-176c30-30.1 8.7-81.9-34-81.9zM224 432L48 256h121.8V80h108.3v176H400L224 432z"/></svg>`,
    primary: true,
  });
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
</Layout>
