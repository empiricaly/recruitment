<script>
  export let group = "";
  export let checked = false;
</script>

<div class="flex items-center">
  <input
    bind:group
    bind:checked
    type="checkbox"
    class="form-checkbox h-4 w-4 text-mint-600 transition duration-150
    ease-in-out" />
  <div class="ml-2">
    <slot />
  </div>
</div>
