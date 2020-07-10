<script>
  import { manageFocus } from "./focus.js";

  export let type = "button";
  export let text;
</script>

<button
  {type}
  class="w-full px-2 py-1 rounded focus:outline-none focus:shadow-outline
  text-gray-100 bg-blue-500 hover:bg-blue-600 active:bg-blue-700
  active:shadow-inner active:border-transparent font-semibold">
  {text}
</button>

<style>
  :global(.focus-disabled) button {
    box-shadow: 0 0 0 3px rgba(66, 153, 225, 0) !important;
  }
</style>
