<script>
  export let color = "gray";
  export let text = "";
  export let large = false;
  export let icon = null;

  $: css = large
    ? `px-3 py-0.5 text-sm  font-semibold leading-5`
    : `px-2.5 py-0.5 text-xs font-semibold leading-4`;
</script>

<span
  class="{css} inline-flex items-center rounded-full bg-{color}-100 text-{color}-800">
  {#if icon}
    <svg
      class="flex-shrink-0 mr-1 h-3 w-3 opacity-75"
      viewBox="0 0 24 24"
      fill="currentColor">
      {@html icon}
    </svg>
  {/if}
  {text}
</span>
