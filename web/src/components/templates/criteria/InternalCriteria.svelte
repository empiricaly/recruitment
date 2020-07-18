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

  export let criteria = {};

  let match = criteria.and ? "and" : "or";
  let prevMatch = match;

  $: {
    if (prevMatch !== match) {
      console.log("new value:", match);
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

  $: console.log(JSON.stringify(criteria, null, "  "));
</script>

<div class="flex items-center mb-3">
  <p class="mr-2">Participants matching</p>
  <Select
    className="w-20"
    id="matchType"
    bind:value={match}
    options={matchTypes}
    placeholder="Matching" />
  <p class="ml-2">of the following conditions:</p>
</div>
<ul>
  <Criteria bind:criteria first />
</ul>
