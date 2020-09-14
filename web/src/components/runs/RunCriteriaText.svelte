<script context="module">
  function getValue(value) {
    if (value.int !== undefined) {
      return value.int;
    } else if (value.float !== undefined) {
      return value.float;
    } else if (value.string !== undefined) {
      return value.string;
    } else if (value.boolean !== undefined) {
      return value.boolean;
    }
  }
</script>

<script>
  import { comparatorsIndex } from "../templates/criteria/criteria.js";

  export let condition = {};
  export let first = false;

  $: operator = condition && !condition.key && (condition.and ? "and" : "or");
  $: console.log(condition, operator, condition[operator]);
</script>

{#if condition}
  {#if operator}
    {#if !first}({/if}
    {#each condition[operator] as c, i (c)}
      {#if i !== 0}{operator}{/if}
      <svelte:self condition={c} />
    {/each}
    {#if !first}){/if}
  {:else}
    {condition.key} is {comparatorsIndex[condition.comparator]}
    {condition.values.map(v => getValue(v)).join(', ')} {' '}
  {/if}
{/if}
