<script>
  import { onMount } from "svelte";

  export let type = "text";
  export let placeholder = "";
  export let focus = false;
  export let value = "";
  export let required = "";

  function init(el) {
    el.type = type;
    if (focus) {
      el.focus();
    }
  }
</script>

<input
  {placeholder}
  {required}
  use:init
  bind:value
  class="w-full px-2 py-1 rounded border focus:outline-none focus:shadow-outline" />
