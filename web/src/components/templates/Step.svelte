<script context="module">
  const stepTypeName = {
    MTURK_HIT: "MTurk HIT Step",
    MTURK_MESSAGE: "MTurk Message Step",
    PARTICIPANT_FILTER: "Participant Filter Step",
  };
  const stepTypeDesc = {
    MTURK_HIT: "TODO: Add Description",
    MTURK_MESSAGE: "TODO: Add Description",
    PARTICIPANT_FILTER: "TODO: Add Description",
  };
</script>

<script>
  import { createEventDispatcher } from "svelte";
  import { uniqueID } from "../../utils/uniq.js";
  import Input from "../base/Input.svelte";
  import Label from "../base/Label.svelte";
  import { confirm } from "../overlays/Alert.svelte";
  import StepMTurkHit from "./StepMTurkHit.svelte";
  import StepMTurkMessage from "./StepMTurkMessage.svelte";
  import StepParticipantFilter from "./StepParticipantFilter.svelte";
  import TemplateSection from "./TemplateSection.svelte";

  export let step;

  const dispatch = createEventDispatcher();

  const uniq = uniqueID();

  async function handleDelete() {
    try {
      await confirm({
        title: "Are you sure?",
        body:
          "You are about to delete the current Step. This cannot be undone.",
        button: "Delete Step",
      });
    } catch (error) {
      return;
    }
    dispatch("delete", { step });
  }
</script>

<TemplateSection title={stepTypeName[step.type]} header>
  <div slot="header">
    <div class="md:grid grid-cols-3 gap-6">
      <div class="flex items-baseline">
        <div class="mr-2 whitespace-no-wrap">
          <Label
            forID={uniq('duration')}
            text="Duration"
            white
            question="Duration of Step in minutes" />
        </div>
        <Input
          id={uniq('duration')}
          type="number"
          min="0"
          right="minutes"
          bind:value={step.duration}
          inputmode="numeric"
          required
          placeholder="0" />
      </div>

      <div />
      <div class="flex items-center md:justify-end">
        <button
          on:click={handleDelete}
          class="flex items-center mt-2 md:mt-0 focus:outline-none">
          <svg
            class="text-gray-50"
            fill="currentColor"
            width="24"
            height="24"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 448 512">
            <path
              d="M440 64H336l-33.6-44.8A48 48 0 0 0 264 0h-80a48 48 0 0 0-38.4
              19.2L112 64H8a8 8 0 0 0-8 8v16a8 8 0 0 0 8 8h18.9l33.2 372.3a48 48
              0 0 0 47.8 43.7h232.2a48 48 0 0 0 47.8-43.7L421.1 96H440a8 8 0 0 0
              8-8V72a8 8 0 0 0-8-8zM171.2 38.4A16.1 16.1 0 0 1 184 32h80a16.1
              16.1 0 0 1 12.8 6.4L296 64H152zm184.8 427a15.91 15.91 0 0 1-15.9
              14.6H107.9A15.91 15.91 0 0 1 92 465.4L59 96h330z" />
          </svg>
          <span class="ml-2 md:hidden">Delete Step</span>
        </button>
      </div>
    </div>
  </div>
  <div slot="description">{stepTypeDesc[step.type]}</div>
  {#if step.type === 'MTURK_HIT'}
    <StepMTurkHit bind:step />
  {:else if step.type === 'MTURK_MESSAGE'}
    <StepMTurkMessage bind:step />
  {:else if step.type === 'PARTICIPANT_FILTER'}
    <StepParticipantFilter bind:step />
  {:else}Unknown Step Type?!?{/if}
</TemplateSection>
