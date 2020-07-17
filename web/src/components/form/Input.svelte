<script>
  import { onMount } from "svelte";

  export let autocomplete = "off";
  export let type = "text";
  export let id = "";
  export let placeholder = "";
  export let focus = false;
  export let value = "";
  export let required = "";
  export let onKeyUp = undefined;
  export let onBlur = undefined;
  export let disabled = false;
  function init(el) {
    el.type = type;
    if (focus) {
      el.focus();
    }

    if (id) {
      el.id = id;
    }
  }
</script>

<div class="mt-1 rounded-md shadow-sm">
  <input
    {autocomplete}
    {placeholder}
    {required}
    {disabled}
    on:keyup={onKeyUp}
    on:blur={onBlur}
    use:init
    bind:value
    class="appearance-none block w-full px-3 py-2 border border-gray-300
    rounded-md placeholder-gray-400 focus:outline-none focus:shadow-outline-blue
    focus:border-blue-300 transition duration-150 ease-in-out sm:text-sm
    sm:leading-5" />
</div>
