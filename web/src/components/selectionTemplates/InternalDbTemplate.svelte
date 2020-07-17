<script context="module">
  const matchTypes = [
    {
      label: "All",
      value: "all"
    },
    {
      label: "Any",
      value: "any"
    }
  ];
</script>

<script>
  import Criteria from "./Criteria.svelte";
  import Select from "../form/Select.svelte";

  export let criteria = {};

  let match = "all";
</script>

<div class="flex items-center mb-3">
  <p class="mr-2">Participants matching</p>
  <Select
    id="matchType"
    bind:value={match}
    options={matchTypes}
    placeholder="Matching" />
  <p class="ml-2">of the following conditions:</p>
</div>
<ul>
  <Criteria bind:criteria />
</ul>
