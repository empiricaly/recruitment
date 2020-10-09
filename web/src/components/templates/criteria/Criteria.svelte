<script>
  import { createEventDispatcher } from "svelte";
  import Button from "../../base/Button.svelte";
  import Input from "../../base/Input.svelte";
  import Select from "../../base/Select.svelte";
  import { comparators } from "./criteria.js";
  import ValueInput from "./ValueInput.svelte";

  const dispatch = createEventDispatcher();

  export let first = false;
  export let criteria = {};

  $: operator = criteria.and && criteria.and.length > 0 ? "and" : "or";

  $: console.log(criteria[operator], operator);

  function handleChildAdd(event) {
    const {
      criteria: child,
      operator: childOperator,
      newCriteria = {
        key: "",
        comparator: "EQUAL_TO",
        values: [],
      },
    } = event.detail;

    const childIndex = criteria[operator].findIndex((c) => c === child);

    if (operator === childOperator) {
      criteria[operator].splice(childIndex + 1, 0, newCriteria);
    } else {
      criteria[operator].splice(childIndex, 1, {
        [childOperator]: [child, newCriteria],
      });
    }
    criteria[operator] = criteria[operator];
  }

  function handleChildRemove(event) {
    const { criteria: child } = event.detail;

    criteria[operator] = criteria[operator].filter((item) => item !== child);
    if (first && criteria[operator].length === 0) {
      criteria[operator] = [
        {
          key: "",
          comparator: "EQUAL_TO",
          values: [],
        },
      ];
    }
    if (criteria[operator].length === 1) {
      dispatch("add", {
        criteria,
        newCriteria: criteria[operator][0],
        operator: operator === "and" ? "or" : "and",
      });
      dispatch("remove", { criteria });
    }
  }

  function handleComparatorChange(event) {
    const { value } = event.detail;
    if (value === "DoesNotExist" || value === "Exists") {
      delete criteria.values;
    } else {
      if (!criteria.values) {
        criteria.values = [];
        criteria = criteria;
      }
    }
  }

  console.log(first);
</script>

<li>
  {#if (!criteria[operator] || criteria[operator].length === 0) && !first}
    <div
      class="min-w-0 flex-1 md:grid md:grid-cols-3 md:gap-2 mt-4 md:mt-1
      items-center">
      <div class="col-span-2 grid grid-cols-3 gap-2 items-center">
        <div>
          <Input bind:value={criteria.key} placeholder="key" />
        </div>
        <div>
          <Select
            bind:value={criteria.comparator}
            options={comparators}
            on:change={handleComparatorChange}
            placeholder="comparator" />
        </div>
        <div>
          <ValueInput bind:criteria />
        </div>
      </div>

      <div class="mt-2 md:mt-0 col-span-1 grid grid-cols-3 gap-2 items-center">
        <Button
          on:click={() => dispatch('add', { criteria, operator: 'and' })}
          text="And"
          secondary />
        <Button
          on:click={() => dispatch('add', { criteria, operator: 'or' })}
          text="Or"
          secondary />
        <Button
          on:click={() => dispatch('remove', { criteria })}
          icon={`<path d="M12 10.586l4.95-4.95 1.414 1.414-4.95 4.95 4.95 4.95-1.414 1.414-4.95-4.95-4.95 4.95-1.414-1.414 4.95-4.95-4.95-4.95L7.05 5.636z"/>`}
          secondary />
      </div>
    </div>
  {:else}
    <div class="relative">
      {#if !first}
        <div class="absolute w-5 h-full pt-3 pb-1 flex justify-center">
          <div class="w-0 h-full border-l border-gray-300 bg-pink-600 relative">
            <div class="absolute inset-0 flex justify-center items-center">
              <div class="bg-white text-gray-400 text-sm">
                {#if criteria.or}OR{:else}AND{/if}
              </div>
            </div>
          </div>
        </div>
      {/if}
      <div class={!first && 'ml-8'}>
        {#each criteria[operator] as c (c)}
          <svelte:self
            bind:criteria={c}
            on:remove={handleChildRemove}
            on:add={handleChildAdd} />
        {/each}
      </div>
    </div>
  {/if}
</li>
