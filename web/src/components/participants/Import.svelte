<script>
  import { mutate } from "svelte-apollo";

  import { notify } from "../../components/overlays/Notification.svelte";
  import Input from "../../components/base/Input.svelte";
  import Label from "../../components/base/Label.svelte";
  import Modal from "../../components/overlays/Modal.svelte";
  import {
    getParticipants,
    setValue,
  } from "../../lib/models/participants/participants.js";
  import { ADD_PARTICIPANTS } from "../../lib/queries";
  import { client } from "../../lib/apollo";
  import { handleErrorMessage } from "../../utils/errorQuery.js";

  export let projectID = undefined;
  export let isOpen = false;

  let files = [];
  let customKey;
  let customValue;

  function handleCancelImport() {
    isOpen = false;
    files = [];
    customKey = null;
    customValue = null;
  }

  function handleImport() {
    let file = files.length > 0 ? files[0] : null;
    let isCustomEmpty;
    let customData = {};

    if (!file) {
      notify({
        failed: true,
        title: `Could not import participants.`,
        body: "No file selected.",
      });
      return;
    }

    if (customKey || customValue) {
      if (
        !customKey ||
        !customKey.trim() ||
        !customValue ||
        !customValue.trim()
      ) {
        isCustomEmpty = true;
      }

      if (isCustomEmpty) {
        notify({
          failed: true,
          title: `Could not import participants.`,
          body: "Custom key/value pair can't be empty.",
        });
        return;
      }

      customData = { [customKey]: setValue(customValue) };
    }

    getParticipants(file, customData, async (newParticipants, error) => {
      if (error) {
        notify({
          failed: true,
          title: `Could not import participants.`,
          body: error,
        });
        return;
      }

      try {
        await mutate(client, {
          mutation: ADD_PARTICIPANTS,
          variables: {
            input: {
              participants: newParticipants,
              projectID,
            },
          },
        });

        files = [];
        isOpen = false;
        customKey = null;
        customValue = null;
        notify({
          failed: false,
          title: `Participants imported.`,
        });
        setTimeout(() => {
          location.reload();
        }, 2000);
      } catch (error) {
        console.log("error", error);
        handleErrorMessage(error);
        notify({
          failed: true,
          title: `Could not import participants`,
          body:
            "Something happened on the server, and we could not import the participants.",
        });
      }
    });
  }
</script>

<Modal
  title="Import Participants"
  button="Import"
  bind:open={isOpen}
  handleAccept={handleImport}
  handleCancel={handleCancelImport}>
  <div class="sm:flex flex-row sm:items-start">
    <input type="file" bind:files />
  </div>
  <div class="mt-4">
    <Label
      text="Custom Data"
      question="Custom data (key value pair) that will be created on imported partcipants." />
  </div>
  <div class="sm:flex flex-row sm:items-start">
    <div class="mt-1">
      <Input bind:value={customKey} placeholder="key" />
    </div>
    <div class="mt-1 ml-3">
      <Input bind:value={customValue} placeholder="value" />
    </div>
  </div>
</Modal>
