<script>
  import { mutate } from "svelte-apollo";
  import { client } from "../../lib/apollo";
  import { addDirtyObject, removeDirtyObject } from "../../lib/dirty";
  import { UPDATE_TEMPLATE } from "../../lib/queries.js";
  import { deepCopy } from "../../utils/copy";
  import { handleErrorMessage } from "../../utils/errorQuery";
  import { debounce } from "../../utils/timing";
  import Button from "../base/Button.svelte";
  import Input from "../base/Input.svelte";
  import Label from "../base/Label.svelte";
  import Select from "../base/Select.svelte";
  import Toggle from "../base/Toggle.svelte";
  import Checkbox from "../base/Checkbox.svelte";
  import { notify } from "../overlays/Notification.svelte";
  import InternalCriteria from "./criteria/InternalCriteria.svelte";
  import MTurkQualifications from "./criteria/MTurkQualifications.svelte";
  import Step from "./Step.svelte";
  import {
    defaultFilterStepArgs,
    defaultHITMessageStepArgs,
    defaultHITStepArgs,
    defaultMessageStepArgs,
    selectionTypes,
  } from "./templates.js";
  import TemplateSection from "./TemplateSection.svelte";

  export let project;
  export let run;
  export let template;
  export let isTemplateDirty;

  let previousTemplate = JSON.stringify(template);
  $: {
    const newTemplate = JSON.stringify(template);
    if (newTemplate !== previousTemplate) {
      isTemplateDirty = true;
      addDirtyObject(template.id);
      save();
    }
  }

  const save = debounce(
    async () => {
      const newTemplate = JSON.stringify(template);
      if (newTemplate === previousTemplate) {
        return;
      }
      previousTemplate = newTemplate;
      console.log("project", project);
      console.log("run", run);
      try {
        const input = {
          runID: run.id,
          template,
        };

        console.log(JSON.stringify(input, null, "  "));

        await mutate(client, {
          mutation: UPDATE_TEMPLATE,
          variables: {
            input,
          },
        });
        isTemplateDirty = false;
        removeDirtyObject(template.id);
      } catch (error) {
        handleErrorMessage(error);
        notify({
          failed: true,
          title: `Could not save Template update`,
          body:
            "Something happened on the server, and we could not save the latest changes to this Run.",
        });
      }
    },
    1000,
    5000
  );

  function handleNewStep(stepType) {
    return function (event) {
      const msgArgs =
        stepType === "MTURK_HIT"
          ? deepCopy(defaultHITMessageStepArgs)
          : deepCopy(defaultMessageStepArgs);
      let params = {
        type: stepType,
        duration: stepType === "MTURK_HIT" || stepType === "WAIT" ? 60 : 0,
        index: template.steps.length,
        msgArgs,
        hitArgs: deepCopy(defaultHITStepArgs),
        filterArgs: deepCopy(defaultFilterStepArgs),
      };

      template.steps[template.steps.length] = params;
    };
  }

  function handleDeleteStep(event) {
    const { step } = event.detail;
    template.steps = template.steps.filter((s) => s !== step);
    for (let i = 0; i < template.steps.length; i++) {
      template.steps[i].index = i;
    }
  }

  function handleMoveStep(event) {
    const { step, isUpward } = event.detail;
    let otherStep;

    template.steps = template.steps.filter((s) => {
      let foundOtherStep = false;

      if (isUpward && step.index - 1 === s.index) {
        s.index = s.index + 1;
        foundOtherStep = true;
      } else if (!isUpward && step.index + 1 === s.index) {
        s.index = s.index - 1;
        foundOtherStep = true;
      }

      if (foundOtherStep) {
        otherStep = s;
        return false;
      }

      return s !== step;
    });

    step.index = isUpward ? step.index - 1 : step.index + 1;

    template.steps = template.steps
      .concat([step, otherStep])
      .sort((a, b) => a.index - b.index);
  }
</script>

<TemplateSection footer>
  <div slot="title">Participant Selection</div>

  <div slot="description">
    A Template starts with the selection of Participants you want in your
    experiment. Internal Database selection uses Participant information
    collected from previous experiments or recruitment runs. MTurk
    Qualifications uses Worker Qualifications defined on MTurk.
  </div>

  <div class="md:grid grid-cols-3 gap-2">
    <div class="max-w-xs">
      <Label forID="selectionType" text="Participant Selection Type" />
      <Select
        id="selectionType"
        bind:value={template.selectionType}
        options={selectionTypes}
        placeholder="Particpant Selection Type" />
    </div>
    {#if template.selectionType === 'INTERNAL_DB'}
      <div class="flex flex-column pt-5 max-w-xs">
        <Checkbox bind:checked={template.internalCriteria.uninitialized}>
          <Label
            forID="unitializedParticipants"
            text="Select Uninitialized Participants"
            question="Uninitialized Participants are Participants that have been imported, and do not yet have a HIT associated with them, 
            so we cannot send messages or bonuses to them. Uninitialized participants are not included in an 
            Internal DB selection. On the contrary, if you check this box, only uninitialized participants 
            will be selected" />
        </Checkbox>
      </div>
    {/if}
  </div>

  <div class="mt-5">
    {#if template.selectionType === 'INTERNAL_DB'}
      <InternalCriteria
        bind:all={template.internalCriteria.all}
        bind:criteria={template.internalCriteria.condition} />
    {:else if template.selectionType === 'MTURK_QUALIFICATIONS'}
      <MTurkQualifications
        sandbox={template.sandbox}
        bind:qualifications={template.mturkCriteria.qualifications} />
    {:else}Unknow Particpant Selection Type{/if}
  </div>

  <div slot="footer">
    <div class="md:grid grid-cols-3 gap-6">
      <div>
        <Label
          forID="participantCount"
          text="Starting Number of Participants"
          question="The number of participants the process should start with
          from the selection phase." />
        <div class="w-28">
          <Input
            id="participantCount"
            type="number"
            min="0"
            bind:value={template.participantCount}
            inputmode="numeric"
            required
            placeholder="0" />
        </div>
      </div>

      <div class="mt-4 md:mt-0">
        <Label
          forID="adult"
          text="May contain Explicit Content"
          question="This run may contain potentially explicit or offensive
          content, for example, nudity." />
        <div class="">
          <Toggle id="adult" bind:checked={template.adult} />
        </div>
      </div>

      <div class="mt-4 md:mt-0">
        <Label
          forID="sandbox"
          text="Use Sandbox"
          question="Use MTurk Sandbox mode instead of production mode. Real
          Workers will not see the HITs. Search for MTurk Sandbox on Google to
          find out more." />
        <div class="">
          <Toggle id="adult" bind:checked={template.sandbox} />
        </div>
      </div>
    </div>
  </div>
</TemplateSection>

<div class="mt-4 hidden sm:block">
  <div class="py-5">
    <div class="border-t border-gray-200" />
  </div>
</div>

{#each template.steps as step}
  <Step
    bind:step
    stepLength={template.steps.length}
    on:moveStep={handleMoveStep}
    on:delete={handleDeleteStep}
    error={template.selectionType === 'MTURK_QUALIFICATIONS' && step.index === 0 && step.type !== 'MTURK_HIT' ? 'First step of a Run using MTurk Qualifications must be an MTurk Hit.' : ''} />
{/each}

{#if template.steps.length > 0}
  <div class="mt-4 hidden sm:block">
    <div class="py-5">
      <div class="border-t border-gray-200" />
    </div>
  </div>
{/if}

<div class="mt-4">
  <h3 class="text-lg font-medium leading-6 text-gray-900">Add Step</h3>
  <p class="mt-1 text-sm leading-5 text-gray-600">
    Steps execute in order a pre-defined time in a Run.
    <!-- TODO: Add Step documentation here. -->
  </p>
</div>

<div class="mt-4 md:grid md:grid-cols-4 md:gap-3">
  <Button
    secondary={template.steps.length > 0}
    on:click={handleNewStep('MTURK_HIT')}
    text="Add MTurk HIT Step" />

  <div class="mt-2 md:mt-0">
    <Button
      secondary={template.steps.length > 0}
      on:click={handleNewStep('MTURK_MESSAGE')}
      disabled={template.steps.length === 0 && template.selectionType !== 'INTERNAL_DB'}
      text="Add MTurk Message Step" />
  </div>

  <div class="mt-2 md:mt-0">
    <Button
      secondary={template.steps.length > 0}
      on:click={handleNewStep('PARTICIPANT_FILTER')}
      disabled={template.steps.length === 0 && template.selectionType !== 'INTERNAL_DB'}
      text="Add Participant Filter Step" />
  </div>

  <div class="mt-2 md:mt-0">
    <Button
      secondary={template.steps.length > 0}
      on:click={handleNewStep('WAIT')}
      disabled={template.steps.length === 0 && template.selectionType !== 'INTERNAL_DB'}
      text="Add Wait Step" />
  </div>
</div>
