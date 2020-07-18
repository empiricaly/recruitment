<script>
  import Button from "../components/base/Button.svelte";
  export let title = "";
  export let overtitle = "";
  export let action = "";
</script>

<div
  class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex justify-between
  items-baseline">
  <h1 class="text-2xl font-semibold text-gray-900 leading-snug">
    {#if overtitle}
      <div
        class="uppercase tracking-wide font-medium text-gray-500 mr-1 text-sm">
        {overtitle}
      </div>
    {/if}
    {title}
  </h1>
  {#if action}
    <Button on:click text={action} />
  {/if}
  <slot name="actions" />
</div>
