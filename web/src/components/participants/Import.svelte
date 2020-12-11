<script>
  import { createEventDispatcher } from "svelte";

  import Input from "../../components/base/Input.svelte";
  import Label from "../../components/base/Label.svelte";
  import Modal from "../../components/overlays/Modal.svelte";
  import { importParticipants } from "../../lib/models/participants/participants.js";

  export let projectID = undefined;
  export let isOpen = false;

  const dispatch = createEventDispatcher();

  let files = [];
  let customKey;
  let customValue;

  function handleCancelImport() {
    isOpen = false;
    files = [];
    customKey = null;
    customValue = null;
  }

  function setOpen(val) {
    isOpen = val;
  }

  async function handleImport() {
    let worker = new Worker(
      `data:text/javascript,
      onmessage = async function(event){    
        ${await importParticipants({
          files,
          customKey,
          customValue,
          dispatch,
          projectID,
          setOpen,
        })}
      };
      `
    );

    worker.postMessage({});
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
