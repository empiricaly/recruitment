<script context="module">
  const selectionTypes = [
    {
      label: "Internal Database",
      value: "internal_db"
    },
    {
      label: "MTurk Qualifications",
      value: "mturk_qualifications"
    }
  ];

  const defaultMessageStepArgs = {
    url: "",
    message: "\n\n\n\n",
    messageType: "PLAIN",
    lobby: "",
    lobbyType: "PLAIN",
    lobbyExpiration: 0
  };

  const defaultHITStepArgs = {
    title: "",
    description: "",
    keywords: "",
    microbatch: false,
    reward: null,
    timeout: 0,
    duration: 60,
    workersCount: 0
  };

  const defaultFilterStepArgs = {
    js: "",
    filter: ""
  };

  const stepTypeArgs = {
    MTURK_HIT: [defaultHITStepArgs, defaultMessageStepArgs],
    MTURK_MESSAGE: [defaultMessageStepArgs],
    PARTICIPANT_FILTER: [defaultFilterStepArgs]
  };
</script>

<script>
  import Layout from "../layouts/Layout.svelte";
  import Button from "../components/base/Button.svelte";
  import Label from "../components/base/Label.svelte";
  import Input from "../components/base/Input.svelte";
  import Toggle from "../components/base/Toggle.svelte";
  import Select from "../components/base/Select.svelte";
  import InternalCriteria from "../components/templates/criteria/InternalCriteria.svelte";
  import MTurkQualifications from "../components/templates/criteria/MTurkQualifications.svelte";
  import TemplateSection from "../components/templates/TemplateSection.svelte";
  import Step from "../components/templates/Step.svelte";
  import { tick } from "svelte";

  let template = {
    name: "Speed Dating 1",
    selectionType: "internal_db",
    participantCount: null,
    adult: false,
    mTurkCriteria: { qualifications: [] },
    internalCriteria: {
      condition: {
        and: [
          {
            key: "",
            comparator: "EqualTo",
            values: []
          }
        ]
      }
    },
    steps: []
  };

  $: console.log(JSON.stringify(template, "", "  "));

  function handleNewStep(stepType) {
    return function(event) {
      const args = [];
      const argGroups = stepTypeArgs[stepType];
      if (!argGroups) {
        return;
      }

      for (const argGroup of argGroups) {
        args.push({ ...argGroup });
      }

      template.steps[template.steps.length] = {
        type: stepType,
        duration: 60,
        args
      };
    };
  }

  function handleDeleteStep(event) {
    const { step } = event.detail;
    template.steps = template.steps.filter(s => s !== step);
  }
</script>

<Layout title={template.name} overtitle="Template">
  <TemplateSection
    title="Participant Selection"
    description="A Procedure starts with the selection of Participants you want
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
      {#if template.selectionType === 'internal_db'}
        <InternalCriteria bind:criteria={template.internalCriteria.condition} />
      {:else if template.selectionType === 'mturk_qualifications'}
        <MTurkQualifications
          bind:qualifications={template.mTurkCriteria.qualifications} />
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
    <Step bind:step on:delete={handleDeleteStep} />
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
        text="Add MTurk Message Step" />
    </div>

    <div class="mt-2 md:mt-0">
      <Button
        secondary={template.steps.length > 0}
        on:click={handleNewStep('PARTICIPANT_FILTER')}
        text="Add Participant Filter Step" />
    </div>
  </div>
</Layout>
