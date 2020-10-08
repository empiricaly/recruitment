<script>
  export let group = "";
  export let checked = false;
  export let id = "";

  function init(el) {
    if (id) {
      el.id = id;
    }
  }
</script>

<div class="flex items-center">
  <input
    use:init
    bind:group
    bind:checked
    type="checkbox"
    class="form-checkbox h-4 w-4 text-mint-600 transition duration-150
    ease-in-out" />
  <label for={id} class="ml-2 block text-sm leading-5 text-gray-900">
    <slot />
  </label>
</div>
