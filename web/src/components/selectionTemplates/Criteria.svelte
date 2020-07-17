<script context="module">
  const comparators = [
    {
      label: "less than",
      value: "LessThan"
    },
    {
      label: "less than or equal to",
      value: "LessThanOrEqualTo"
    },
    {
      label: "greater than",
      value: "GreaterThan"
    },
    {
      label: "greater than or equal to",
      value: "GreaterThanOrEqualTo"
    },
    {
      label: "equal to",
      value: "EqualTo"
    },
    {
      label: "not equal to",
      value: "NotEqualTo"
    },
    {
      label: "exists",
      value: "Exists"
    },
    {
      label: "does not exist",
      value: "DoesNotExist"
    },
    {
      label: "in",
      value: "In"
    },
    {
      label: "not in",
      value: "NotIn"
    }
  ];
</script>

<script>
  import _ from "lodash";

  import { createEventDispatcher } from "svelte";

  import ValueInput from "./ValueInput.svelte";

  import Select from "../form/Select.svelte";
  import Input from "../form/Input.svelte";
  import Button from "../form/Button.svelte";

  const dispatch = createEventDispatcher();

  export let criteria = {};

  $: operator = !criteria.key && (criteria.and ? "and" : "or");

  function handleChildAdd(event) {
    const {
      criteria: child,
      operator: childOperator,
      newCriteria = {
        key: "",
        comparator: "EqualTo",
        values: []
      }
    } = event.detail;

    const childIndex = criteria[operator].findIndex(c => _.isEqual(c, child));

    if (operator === childOperator) {
      criteria[operator].splice(childIndex + 1, 0, newCriteria);
    } else {
      criteria[operator].splice(childIndex, 1, {
        [childOperator]: [child, newCriteria]
      });
    }
    criteria[operator] = criteria[operator];
  }

  function handleChildRemove(event) {
    const { criteria: child } = event.detail;

    criteria[operator] = criteria[operator].filter(
      item => !_.isEqual(item, child)
    );
    if (criteria[operator].length === 1) {
      dispatch("add", {
        criteria,
        newCriteria: criteria[operator][0],
        operator: operator === "and" ? "or" : "and"
      });
      dispatch("remove", { criteria });
    }
  }

  function handleComparatorChange() {
    if (
      criteria.comparator === "DoesNotExist" ||
      (criteria.comparator === "Exists" && criteria.values)
    ) {
      delete criteria.values;
    } else {
      if (!criteria.values) {
        criteria.values = [];
        criteria = criteria;
      }
    }
  }
</script>

<li>
  {#if criteria.key !== undefined}
    <div class="min-w-0 flex-1 md:grid md:grid-cols-6 md:gap-4 mb-1">
      <Input bind:value={criteria.key} placeholder="key" />
      <Select
        bind:value={criteria.comparator}
        options={comparators}
        on:change={handleComparatorChange}
        placeholder="comparator" />
      <ValueInput bind:criteria />
      <Button
        on:click={() => dispatch('add', { criteria, operator: 'and' })}
        text="And" />
      <Button
        on:click={() => dispatch('add', { criteria, operator: 'or' })}
        text="Or" />
      <Button
        on:click={() => dispatch('remove', { criteria })}
        text="Remove"
        secondary />
    </div>
  {:else}
    <div class="relative ml-4 border-l-2 border-gray-300">
      <span
        class="absolute w-10 h-10 top-1/2 bg-white flex justify-center
        items-center"
        style="left: -20px">
        {#if criteria.or}OR{:else}AND{/if}
      </span>
      <div class="ml-8">
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
