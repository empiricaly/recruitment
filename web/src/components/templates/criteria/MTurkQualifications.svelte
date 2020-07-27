<script>
  import MTurkCriteria from "./MTurkCriteria.svelte";
  import Select from "../../base/Select.svelte";
  import Button from "../../base/Button.svelte";
  import Radio from "../../base/Radio.svelte";
  import Checkbox from "../../base/Checkbox.svelte";

  export let qualifications = [];
  export let useMaster = false;

  const maxQualification = 5;

  let useMasterValue = useMaster ? "useMaster" : "doNotUseMaster";

  $: {
    if (useMasterValue) {
      useMaster = useMasterValue === "useMaster";
    }
  }

  function handleQualRemove(event) {
    const { index } = event.detail;
    qualifications.splice(index, 1);
    qualifications = qualifications;
  }

  function handleChange(event) {
    const { value } = event.detail;
  }
</script>

<div class="font-bold text-gray-600 text-sm tracking-wide">
  Require that Workers be Masters to do your tasks
</div>

<div class="mt-2 h-8 flex items-center">
  <Radio bind:group={useMasterValue} value="useMaster">
    <div class="flex items-center">Yes</div>
  </Radio>
  <div class="ml-4">
    <Radio bind:group={useMasterValue} value="doNotUseMaster">
      <div class="flex items-center">No</div>
    </Radio>
  </div>
</div>

<div class="font-bold text-gray-600 text-sm tracking-wide mt-5">
  Specify any additional qualifications Workers must meet to work on your tasks:
</div>

<ul class="mt-2">
  {#each qualifications as q, i (q)}
    <MTurkCriteria
      bind:qualification={q}
      on:remove={handleQualRemove}
      index={i} />
  {/each}
</ul>

<div class="text-xs w-80 mt-2 flex items-center">
  {#if qualifications.length < maxQualification}
    <Button
      secondary
      on:click={() => {
        qualifications = [...qualifications, { id: '', comparator: '', values: [], locales: [] }];
      }}
      text="Add Another Criterion" />
    <span class="ml-3 text-orange-600">
      up to {qualifications.length === 0 ? maxQualification : maxQualification - qualifications.length + ' more'}
    </span>
  {/if}
</div>
