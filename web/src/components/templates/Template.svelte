<script>
  import { mutate } from "svelte-apollo";
  import { client } from "../../lib/apollo";
  import { UPDATE_TEMPLATE } from "../../lib/queries.js";
  import { deepCopy } from "../../utils/copy";
  import { debounce } from "../../utils/timing";
  import Button from "../base/Button.svelte";
  import Input from "../base/Input.svelte";
  import Label from "../base/Label.svelte";
  import Select from "../base/Select.svelte";
  import Toggle from "../base/Toggle.svelte";
  import { notify } from "../overlays/Notification.svelte";
  import InternalCriteria from "./criteria/InternalCriteria.svelte";
  import MTurkQualifications from "./criteria/MTurkQualifications.svelte";
  import Step from "./Step.svelte";
  import {
    defaultFilterStepArgs,
    defaultHITStepArgs,
    defaultMessageStepArgs,
    selectionTypes,
  } from "./templates.js";
  import TemplateSection from "./TemplateSection.svelte";

  export let project;
  export let run;
  export let template;

  // $: console.log(JSON.stringify(template, "", "  "));
  // $: console.log(JSON.stringify(template));

  $: {
    if (template) {
      save();
    }
  }

  let previousTemplate = JSON.stringify(template);
  const save = debounce(
    async () => {
      const newTemplate = JSON.stringify(template);
      if (newTemplate === previousTemplate) {
        console.log("nothing changed");
        return;
      }
      previousTemplate = newTemplate;
      console.log("project", project);
      console.log("run", run);
      try {
        const input = {
          runID: run.id,
          projectID: project.id,
          template,
        };

        console.log(JSON.stringify(input, null, "  "));

        await mutate(client, {
          mutation: UPDATE_TEMPLATE,
          variables: {
            input,
          },
        });
        notify({
          success: true,
          title: `Run Saved`,
        });
      } catch (error) {
        console.error(error);
        notify({
          failed: true,
          title: `Could not save Template update`,
          body:
            "Something happened on the server, and we could not save the latest changes to this Run.",
        });
      }
    },
    2500,
    30000
  );

  function handleNewStep(stepType) {
    return function (event) {
      let params = {
        type: stepType,
        duration: 60,
        index: template.steps.length,
        msgArgs: deepCopy(defaultMessageStepArgs),
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

<TemplateSection
  title="Participant Selection"
  description="A Template starts with the selection of Participants you want
    in your experiment. Internal Database selection uses Participant information
    collected from previous experiments or recruitment runs. MTurk
    Qualifications uses Worker Qualifications defined on MTurk."
  footer>
  <div class="max-w-xs">
    <Label forID="selectionType" text="Participant Selection Type" />
    <Select
      id="selectionType"
      bind:value={template.selectionType}
      options={selectionTypes}
      placeholder="Particpant Selection Type" />
  </div>

  <div class="mt-5">
    {#if template.selectionType === 'INTERNAL_DB'}
      <InternalCriteria
        bind:all={template.internalCriteria.all}
        bind:criteria={template.internalCriteria.condition} />
    {:else if template.selectionType === 'MTURK_QUALIFICATIONS'}
      <MTurkQualifications
        bind:qualifications={template.mturkCriteria.qualifications} />
    {:else}Unknow Particpant Selection Type{/if}
  </div>

  <div slot="footer">
    <div class="md:grid grid-cols-2 gap-6">
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

      <div class="flex items-center mt-4 md:mt-0">
        <Label
          forID="adult"
          text="This project may contain potentially explicit or offensive
            content, for example, nudity." />
        <div class="ml-8">
          <Toggle id="adult" bind:checked={template.adult} />
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
    on:delete={handleDeleteStep} />
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
    Steps execute in order a pre-defined time in a Run. TODO: Add Step
    documentation here.
  </p>
</div>

<div class="mt-4 md:grid md:grid-cols-3 md:gap-4">
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
</div>
