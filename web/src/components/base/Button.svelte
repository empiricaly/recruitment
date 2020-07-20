<script>
  export let type = "button";
  export let text = "";
  export let icon = "";
  export let secondary = false;
  export let full = false;

  $: css = secondary
    ? "text-indigo-700 bg-indigo-100 hover:bg-indigo-50 focus:border-indigo-300 focus:shadow-outline-indigo active:bg-indigo-200"
    : "text-white bg-indigo-600 hover:bg-indigo-500 focus:border-indigo-700 focus:shadow-outline-indigo active:bg-indigo-700";
</script>

<span class:w-full={full} class="block rounded-md shadow-sm">
  <button
    on:click
    {type}
    class="{css} flex w-full justify-center px-4 py-2 border border-transparent
    text-sm font-medium rounded-md transition duration-150 ease-in-out
    focus:outline-none">
    {#if icon}
      <svg
        class="h-5 w-5 min-w-6 min-h-6 flex-shrink-0 "
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
