<script>
  import { query } from "svelte-apollo";

  import ParticipantList from "../components/participants/ParticipantList.svelte";
  import Import from "../components/participants/Import.svelte";

  import Header from "../layouts/Header.svelte";
  import { client } from "../lib/apollo";
  import { GET_ALL_PARTICIPANTS } from "../lib/queries";
  import {
    participantPerQueryType,
    exportCSV,
    exportJson,
  } from "../lib/models/participants/participants.js";
  import { handleErrorMessage } from "../utils/errorQuery.js";

  let keys;
  let participants;
  let isOpen = false;

  async function fetchAllParticipants() {
    let args = {
      query: GET_ALL_PARTICIPANTS,
    };
    const participantsQuery = query(client, args);

    try {
      const result = await participantsQuery.refetch();
      const pp = participantPerQueryType("all", result);
      if (pp) {
        participants = pp.participants;
      }
    } catch (error) {
      handleErrorMessage(error);
    }
  }

  async function handleClick(event) {
    const { action } = event.detail;
    switch (action) {
      case "exportjson": {
        await fetchAllParticipants();
        exportJson(participants, keys);
        break;
      }
      case "exportcsv": {
        await fetchAllParticipants();
        exportCSV(participants, keys);
        break;
      }
      case "import":
        isOpen = true;
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
    text: "Export JSON",
    action: "exportjson",
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M400 208h-73.8V80c0-26.5-21.5-48-48-48H169.8c-26.5 0-48 21.5-48 48v128H48.1c-42.6 0-64.2 51.7-33.9 81.9l175.9 176c18.7 18.7 49.1 18.7 67.9 0l176-176c30-30.1 8.7-81.9-34-81.9zM224 432L48 256h121.8V80h108.3v176H400L224 432z"/></svg>`,
    primary: true,
  });
</script>

<div class="pt-2 pb-6 md:py-6">
  <Header title="All Participants" on:click={handleClick} {actions} />

  <div class="py-4">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
      <ParticipantList
        type="all"
        queryArgs={{ query: GET_ALL_PARTICIPANTS }}
        bind:keys />
    </div>
  </div>
  <Import bind:isOpen />
</div>
