<script>
  export let group;
  export let value;
  export let name = value;
</script>

<label class="inline-flex items-center">
  <input
    bind:group
    {value}
    type="radio"
    class="form-radio text-indigo-600"
    {name} />
  <span class="ml-2">
    <slot />
  </span>
</label>
