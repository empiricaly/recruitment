<script>
  export let type = "button";
  export let text = "";
  export let icon = "";
  export let secondary = false;
  export let tertiary = false;
  export let full = false;
  export let disabled = false;

  $: css = disabled
    ? "text-gray-400 bg-gray-200 hover:bg-gray-50 focus:border-gray-300 active:bg-gray-200"
    : tertiary
    ? "border-gray-300 text-gray-700 bg-white hover:text-gray-500 focus:border-mint-300 active:text-gray-800 active:bg-gray-50 active:text-gray-800"
    : secondary
    ? "text-mint-700 bg-mint-100 hover:bg-mint-50 focus:border-mint-300 active:bg-mint-200"
    : "text-white bg-mint-600 hover:bg-mint-500 focus:border-mint-700 active:bg-mint-700";
</script>

<span class:w-full={full} class="block rounded-md shadow-sm">
  <button
    on:click
    {type}
    {disabled}
    class="{css} flex w-full justify-center items-center px-4 py-2 border
    border-transparent text-sm font-medium rounded-md transition duration-150
    ease-in-out focus:outline-none focus:shadow-outline">
    {#if icon}
      <svg
        class="flex-shrink-0 mr-1.5 h-4 w-4 "
        fill="currentColor"
        viewBox="0 0 24 24">
        {@html icon}
      </svg>
    {/if}
    {text}
  </button>
</span>

<style>
  :global(.focus-disabled) button {
    box-shadow: 0 0 0 3px rgba(66, 153, 225, 0) !important;
  }
</style>
