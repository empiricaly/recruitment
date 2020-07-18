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
</script>

<script>
  import Layout from "../layouts/Layout.svelte";
  import Label from "../components/base/Label.svelte";
  import Select from "../components/base/Select.svelte";
  import InternalDbTemplate from "../components/selectionTemplates/InternalDbTemplate.svelte";

  let template = {
    name: "Speed Dating 1",
    selectionType: "internal_db",
    mTurkCriteria: { qualifications: [] },
    internalCriteria: {
      condition: {
        and: [
          {
            key: "age",
            comparator: "GreaterThan",
            values: [{ int: 18 }]
          },
          {
            key: "country",
            comparator: "In",
            values: [{ string: "us" }, { string: "id" }]
          },
          {
            or: [
              {
                key: "experiment_123",
                comparator: "DoesNotExist"
              },
              {
                key: "experiment_321",
                comparator: "DoesNotExist"
              },
              {
                key: "experiment_343",
                comparator: "DoesNotExist"
              }
            ]
          }
        ]
      }
    },
    steps: []
  };
</script>

<Layout title={template.name} overtitle="Template">
  <!-- These actions are for debug! -->
  <!-- <div slot="actions" class="text-gray-300">
    Debug –
    <button on:click={() => (status = 'running')}>Something</button>
    |
    <button on:click={() => (status = 'running')}>Something else</button>
    – Debug
  </div> -->

  <section class="mt-8 md:grid md:grid-cols-10 md:gap-6">

    <div class="md:mt-0 md:col-span-3">
      <div class="px-4 sm:px-0">
        <h3 class="text-lg font-medium leading-6 text-gray-900">
          Participant Selection
        </h3>
        <p class="mt-1 text-sm leading-5 text-gray-600">
          A Procedure starts with the selection of Participants you want in your
          experiment. Internal Database selection uses Participant information
          collected from previous experiments or recruitment runs. MTurk
          Qualifications uses Worker Qualifications defined on MTurk.
        </p>
      </div>

    </div>

    <div class="mt-5 md:mt-0 md:col-span-7">
      <div class="px-4 py-5 shadow sm:rounded-md bg-white sm:p-6">
        <Label forID="selectionType" text="Participant Selection Type" />
        <Select
          id="selectionType"
          bind:value={template.selectionType}
          options={selectionTypes}
          placeholder="Particpant Selection Type" />

        <div class="mt-5">
          {#if template.selectionType === 'internal_db'}
            <InternalDbTemplate
              bind:criteria={template.internalCriteria.condition} />
          {:else if template.selectionType === 'mturk_qualifications'}
            MTurk
          {:else}Unknow Particpant Selection Type{/if}
        </div>
      </div>
    </div>

  </section>

  <div class="hidden sm:block">
    <div class="py-5">
      <div class="border-t border-gray-200" />
    </div>
  </div>

</Layout>
