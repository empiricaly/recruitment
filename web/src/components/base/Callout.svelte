<script>
  export let color = "gray";
  export let title = null;
  export let icon = null;
</script>

<div class="rounded-md bg-{color}-50 p-4">
  <div class="flex">
    {#if icon}
      <div class="flex-shrink-0">
        <svg
          class="h-5 w-5 text-{color}-400"
          viewBox="0 0 20 20"
          fill="currentColor">
          {@html icon}
        </svg>
      </div>
    {/if}
    <div class={icon ? 'ml-3' : ''}>
      {#if title}
        <h3 class="text-sm leading-5 font-medium text-{color}-800">{title}</h3>
      {/if}
      <div class="{title ? 'mt-2' : ''} text-sm leading-5 text-{color}-700">
        <slot />
      </div>
    </div>
  </div>
</div>
