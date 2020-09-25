<script context="module">
  const bools = [
    { label: "True", value: 1 },
    { label: "False", value: 0 },
  ];
  const comparisons = [
    {
      label: "<",
      value: "LESS_THAN",
      title: "lesser than",
    },
    {
      label: "≤",
      value: "LESS_THAN_OR_EQUAL_TO",
      title: "lesser than or equal to",
    },
    {
      label: ">",
      value: "GREATER_THAN",
      title: "greater than",
    },
    {
      label: "≥",
      value: "GREATER_THAN_OR_EQUAL_TO",
      title: "greater than or equal to",
    },
  ];
  const locations = [
    {
      label: "is",
      value: "EQUAL_TO",
    },
    {
      label: "is not",
      value: "NOT_EQUAL_TO",
    },
    {
      label: "is one of",
      value: "IN",
    },
    {
      label: "is not one of",
      value: "NOT_IN",
    },
  ];
  const customs = [
    {
      label: "Has Been Granted",
      value: "EXISTS",
    },
    {
      label: "Has Not Been Granted",
      value: "DOES_NOT_EXIST",
    },
    ...comparisons,
    ...locations,
  ];

  function mapQualTypes(qualTypes) {
    return qualTypes.map((q) => {
      return { label: q.name, value: q.id };
    });
  }

  function mapLocales(locales) {
    return locales.map(({ country, subdivision }) => {
      const label = subdivision ? `${country}-${subdivision}` : country;
      const value = { country, subdivision };
      return { label, value };
    });
  }

  function mapComparator(selectedQual) {
    switch (selectedQual.type) {
      case "CUSTOM":
        return customs;

      case "COMPARISON":
        return comparisons;

      case "LOCATION":
        return locations;

      default:
        return [];
    }
  }

  function mapIntegers(qualificationId) {
    const integers = [];
    const isHITsApproved = qualificationId === "00000000000000000040";

    // Number of HITs approved return different list of integers
    if (isHITsApproved) {
      const intHits = [0, 50, 100, 500, 1000, 5000, 10000];
      intHits.forEach((h) => {
        integers.push({ label: h, value: h });
      });
    } else {
      for (let index = 0; index <= 100; index++) {
        integers.push({ label: index, value: index });
      }
    }
    return integers;
  }
</script>

<script>
  import { createEventDispatcher } from "svelte";
  import { query } from "svelte-apollo";
  import { client } from "../../../lib/apollo";
  import Select from "../../base/Select.svelte";
  import Input from "../../base/Input.svelte";
  import Button from "../../base/Button.svelte";

  const dispatch = createEventDispatcher();

  import {
    MTURK_LOCALES,
    MTURK_QUALIFICATION_TYPES,
  } from "../../../lib/queries";

  export let qualification = {};
  export let index = undefined;

  let selectedQual;

  function handleIdChange(event, quals) {
    const { value: id } = event.detail;
    qualification.comparator = "";
    qualification.values = [];
    qualification.locales = [];

    selectedQual = quals.find((q) => q.id === id);
  }

  function handleComparatorChange(event) {
    qualification.values = [];
    qualification.locales = [];
  }

  $: isLocation = selectedQual && selectedQual.type === "LOCATION";
  $: isPremium = selectedQual && selectedQual.type === "BOOL";
  $: isPresence =
    qualification &&
    qualification.comparator &&
    (qualification.comparator === "DoesNotExist" ||
      qualification.comparator === "Exists");
  $: isMultiSelect =
    qualification &&
    qualification.comparator &&
    (qualification.comparator === "In" || qualification.comparator === "NotIn");

  $: mturkLocales = query(client, { query: MTURK_LOCALES });
  $: mturkQualTypes = query(client, { query: MTURK_QUALIFICATION_TYPES });
</script>

<li>
  <div
    class="min-w-0 flex-1 md:grid md:grid-cols-3 md:gap-2 mt-4 md:mt-2
      items-center">
    {#await $mturkQualTypes}
      Loading...
    {:then result}
      <div>
        <Select
          bind:value={qualification.id}
          options={mapQualTypes(result.data.mturkQualificationTypes)}
          on:change={(event) => handleIdChange(event, result.data.mturkQualificationTypes)}
          placeholder="Select Qualification" />
      </div>

      {#if qualification.id && !isPremium}
        <Select
          bind:value={qualification.comparator}
          options={mapComparator(selectedQual)}
          on:change={handleComparatorChange}
          placeholder="Select Comparator" />
      {/if}

      <div class="flex">
        {#if qualification.id}
          {#if isPremium}
            <Select
              bind:value={qualification.values}
              options={bools}
              placeholder="Select Value" />
          {/if}

          {#if qualification.comparator && isLocation}
            {#await $mturkLocales then res}
              <Select
                bind:value={qualification.locales}
                options={mapLocales(res.data.mturkLocales)}
                multiple={isMultiSelect}
                placeholder="Select Locales" />
            {/await}
          {/if}

          {#if qualification.comparator && qualification.comparator && !isLocation && !isPremium && !isPresence}
            <div class="w-24">
              <Select
                bind:value={qualification.values}
                options={mapIntegers(qualification.id)}
                multiple={isMultiSelect}
                placeholder="Select Value" />
            </div>
          {/if}
        {/if}
        <div class="ml-2">
          <Button
            on:click={() => dispatch('remove', { index })}
            icon={`<path d="M12 10.586l4.95-4.95 1.414 1.414-4.95 4.95 4.95 4.95-1.414 1.414-4.95-4.95-4.95 4.95-1.414-1.414 4.95-4.95-4.95-4.95L7.05 5.636z"/>`}
            secondary />
        </div>
      </div>
    {:catch error}
      Error loading MTurk qualification types: {error}
    {/await}
  </div>
</li>
