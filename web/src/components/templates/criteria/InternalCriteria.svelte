<script context="module">
  const matchTypes = [
    {
      label: "All",
      value: "and"
    },
    {
      label: "Any",
      value: "or"
    }
  ];
</script>

<script>
  import Criteria from "./Criteria.svelte";
  import Select from "../../base/Select.svelte";
  import Radio from "../../base/Radio.svelte";

  export let criteria = {};
  export let useCriteria = false;
  let useCriteriaValue = useCriteria ? "useCriteria" : "doNotUseCriteria";

  $: {
    if (useCriteriaValue) {
      useCriteria = useCriteriaValue === "useCriteria";
    }
  }

  let match = criteria.and ? "and" : "or";
  let prevMatch = match;

  $: {
    if (prevMatch !== match) {
      const oldcrit = criteria[match === "or" ? "and" : "or"];
      const newCrit = [];
      for (const crit of oldcrit) {
        if (crit[match]) {
          for (const inner of crit[match]) {
            newCrit.push(inner);
          }
        } else {
          newCrit.push(crit);
        }
      }

      const crit = { [match]: newCrit };
      criteria = crit;
      prevMatch = match;
    }
  }

  function handleChange(event) {
    const { value } = event.detail;
  }
</script>

<div class="font-semibold text-gray-400 uppercase text-sm tracking-wide">
  Participant selection
</div>

<div class="mt-2 h-8 flex items-center">
  <Radio bind:group={useCriteriaValue} value="doNotUseCriteria">
    <div class="flex items-center">
      All Participants in the internal database.
    </div>
  </Radio>
</div>

<div class="h-8 flex items-center">
  <Radio bind:group={useCriteriaValue} value="useCriteria">
    {#if useCriteria}
      <div class="flex items-center">
        <p class="mr-2">Participants matching</p>
        <Select
          className="w-20"
          id="matchType"
          bind:value={match}
          options={matchTypes}
          placeholder="Matching" />
        <p class="ml-2">of the following conditions:</p>
      </div>
    {:else}
      <div>Participants matching specific criteria.</div>
    {/if}
  </Radio>
</div>

{#if useCriteria}
  <ul class="mt-3">
    <Criteria bind:criteria first />
  </ul>
{/if}
