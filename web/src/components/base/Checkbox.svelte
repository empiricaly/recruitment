<script>
  export let group;
  export let value;
  export let value;

  const id = "randomid";
</script>

<div class="flex items-center">
  <input
    {id}
    bind:group
    bind:value
    type="checkbox"
    class="form-checkbox h-4 w-4 text-indigo-600 transition duration-150
    ease-in-out" />
  <label for={id} class="ml-2 block text-sm leading-5 text-gray-900">
    <slot />
  </label>
</div>
