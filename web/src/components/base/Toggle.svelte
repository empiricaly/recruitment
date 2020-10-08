<script>
  export let checked;
  export let id = null;

  function init(el) {
    if (id) {
      el.id = id;
    }
  }
</script>

<span
  on:click={() => (checked = !checked)}
  role="checkbox"
  tabindex="0"
  aria-checked={checked}
  class="{checked ? 'bg-mint-600' : 'bg-gray-200'} relative inline-flex
  flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer
  transition-colors ease-in-out duration-200 focus:outline-none
  focus:shadow-outline-mint">
  <span
    aria-hidden="true"
    class="{checked ? 'translate-x-5' : 'translate-x-0'} inline-block h-5 w-5
    rounded-full bg-white shadow transform transition ease-in-out duration-200" />
</span>
