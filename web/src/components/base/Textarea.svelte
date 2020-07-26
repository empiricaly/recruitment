<script>
  export let id = "";
  export let placeholder = "";
  export let rows = "3";
  export let pattern = null;
  export let focus = false;
  export let value = "";
  export let required = "";
  export let disabled = false;

  function init(el) {
    if (focus) {
      el.focus();
    }

    if (id) {
      el.id = id;
    }

    if (pattern) {
      el.pattern = pattern;
    }
  }
</script>

<div class="flex rounded-md shadow-sm">
  <textarea
    id="about"
    {placeholder}
    {rows}
    {required}
    {disabled}
    on:keyup
    on:blur
    on:focus
    use:init
    bind:value
    class="form-textarea block w-full transition duration-150 ease-in-out
    sm:text-sm sm:leading-5" />
</div>
