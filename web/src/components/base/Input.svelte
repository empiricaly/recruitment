<script context="module">
  const inserts = {
    $: 7,
    USD: 12,
    seconds: 20,
    minutes: 20
  };
</script>

<script>
  export let autocomplete = "off";
  export let type = "text";
  export let id = "";
  export let placeholder = "";
  export let pattern = null;
  export let inputmode = null;
  export let min = null;
  export let max = null;
  export let focus = false;
  export let value = "";
  export let required = "";
  export let disabled = false;
  export let left = null;
  export let right = null;

  function init(el) {
    el.type = type;
    if (min !== null) {
      el.min = min;
    }
    if (max !== null) {
      el.max = max;
    }

    if (id) {
      el.id = id;
    }

    if (type === "number" && !pattern) {
      pattern = "d*";
    }
    if (pattern) {
      el.pattern = pattern;
    }

    if (type === "number" && !inputmode) {
      inputmode = "numeric";
    }
    if (inputmode) {
      el.inputmode = inputmode;
    }

    if (focus) {
      el.focus();
    }
  }
</script>

<div class="relative rounded-md shadow-sm">
  {#if left}
    <div
      class="absolute inset-y-0 left-0 pl-3 flex items-center
      pointer-events-none">
      <span class="text-gray-500 sm:text-sm sm:leading-5">{left}</span>
    </div>
  {/if}
  <input
    {autocomplete}
    {placeholder}
    {required}
    {disabled}
    on:keyup
    on:blur
    on:focus
    use:init
    bind:value
    class="appearance-none block w-full pl-{left ? inserts[left] : 3} pr-{right ? inserts[right] : 3}
    text-gray-900 py-2 border border-gray-300 rounded-md placeholder-gray-400
    focus:outline-none focus:shadow-outline-blue focus:border-blue-300
    transition duration-150 ease-in-out sm:text-sm sm:leading-5" />
  {#if right}
    <div
      class="absolute inset-y-0 right-0 pr-3 flex items-center
      pointer-events-none">
      <span class="text-gray-500 sm:text-sm sm:leading-5" id="price-currency">
        {right}
      </span>
    </div>
  {/if}
</div>

<style>
  /* Chrome, Safari, Edge, Opera */
  div :global(input::-webkit-outer-spin-button),
  div :global(input::-webkit-inner-spin-button) {
    -webkit-appearance: none;
    margin: 0;
  }

  /* Firefox */
  div :global(input[type="number"]) {
    -moz-appearance: textfield;
  }
</style>
