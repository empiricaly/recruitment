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
  export let stepLength;
  export let error = "";

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

  function handleMoveStep(isUpward) {
    dispatch("moveStep", { step, isUpward });
  }

  $: isFirstStep = step.index === 0;
  $: isLastStep = step.index + 1 === stepLength;
</script>

<TemplateSection title={stepTypeName[step.type]} header invalid={error !== ''}>
  <div slot="header">
    <div class="flex justify-between">
      <div class="mr-2 flex items-baseline">
        {#if error !== ''}
          <div
            class="bg-red-100 text-red-700 rounded px-4 py-2 flex items-center">
            <svg
              class="flex-shrink-0 h-6 w-6 text-red-500 mr-3"
              xmlns="http://www.w3.org/2000/svg"
              fill="currentColor"
              viewBox="0 0 512 512">
              <path
                d="M256 8C119 8 8 119 8 256s111 248 248 248 248-111
              248-248S393 8 256 8zm0 464c-118.7 0-216-96.1-216-216 0-118.7
              96.1-216 216-216 118.7 0 216 96.1 216 216 0 118.7-96.1
              216-216 216zm94.8-285.3L281.5 256l69.3 69.3c4.7 4.7 4.7 12.3
              0 17l-8.5 8.5c-4.7 4.7-12.3 4.7-17 0L256 281.5l-69.3
              69.3c-4.7 4.7-12.3 4.7-17 0l-8.5-8.5c-4.7-4.7-4.7-12.3
              0-17l69.3-69.3-69.3-69.3c-4.7-4.7-4.7-12.3
              0-17l8.5-8.5c4.7-4.7 12.3-4.7 17 0l69.3 69.3
              69.3-69.3c4.7-4.7 12.3-4.7 17 0l8.5 8.5c4.6 4.7 4.6 12.3 0
              17z" />
            </svg>

            {error}
            {error}
          </div>
        {:else}
          <div class="ml-2 mr-2 whitespace-no-wrap">
            <Label
              forID={uniq('duration')}
              text="Duration"
              white
              question="Duration of Step in minutes" />
          </div>
          <div class="w-32">
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
        {/if}
      </div>

      <div class="ml-2 flex items-center md:justify-end">
        {#if stepLength !== 1 && !isFirstStep}
          <button
            on:click={() => handleMoveStep(true)}
            class="flex items-center mt-2 md:mt-0 focus:outline-none">
            <svg
              class="text-gray-50"
              fill="currentColor"
              width="24"
              height="24"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 448 512">
              <path
                d="M6.101 261.899L25.9 281.698c4.686 4.686 12.284 4.686 16.971 
              0L198 126.568V468c0 6.627 5.373 12 12 12h28c6.627 0 12-5.373 
              12-12V126.568l155.13 155.13c4.686 4.686 12.284 4.686 16.971 0l19.799-19.799c4.686-4.686 
              4.686-12.284 0-16.971L232.485 35.515c-4.686-4.686-12.284-4.686-16.971 0L6.101 244.929c-4.687 
              4.686-4.687 12.284 0 16.97z" />
            </svg>
            <span class="ml-2 md:hidden">Upward Step</span>
          </button>
        {/if}
        {#if stepLength !== 1 && !isLastStep}
          <button
            on:click={() => handleMoveStep(false)}
            class="flex items-center mt-2 ml-4 md:mt-0 focus:outline-none">
            <svg
              class="text-gray-50"
              fill="currentColor"
              width="24"
              height="24"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 448 512">
              <path
                d="M441.9 250.1l-19.8-19.8c-4.7-4.7-12.3-4.7-17 0L250 385.4V44c0-6.6-5.4-12-12-12h-28c-6.6 
              0-12 5.4-12 12v341.4L42.9 230.3c-4.7-4.7-12.3-4.7-17 0L6.1 250.1c-4.7 4.7-4.7 12.3 0 17l209.4 
              209.4c4.7 4.7 12.3 4.7 17 0l209.4-209.4c4.7-4.7 4.7-12.3 0-17z" />
            </svg>
            <span class="ml-2 md:hidden">Downward Step</span>
          </button>
        {/if}
        <button
          on:click={handleDelete}
          class="flex items-center mt-2 ml-4 mr-2 md:mt-0 focus:outline-none">
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
