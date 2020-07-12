<script>
  import { manageFocus } from "./focus.js";

  export let type = "button";
  export let text;
  export let full = false;
</script>

<span class:w-full={full} class="block rounded-md shadow-sm">
  <button
    {type}
    class="flex w-full justify-center py-2 px-4 border border-transparent
    text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-500
    focus:outline-none focus:border-indigo-700 focus:shadow-outline-indigo
    active:bg-indigo-700 transition duration-150 ease-in-out">
    {text}
  </button>
</span>

<style>
  :global(.focus-disabled) button {
    box-shadow: 0 0 0 3px rgba(66, 153, 225, 0) !important;
  }
</style>
