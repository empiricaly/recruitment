<script>
  export let title = "";
  export let description = "";
  export let header = false;
  export let footer = false;
  export let invalid = false;
  export let current = false;
</script>

<section class="mt-4 md:grid md:grid-cols-10 md:gap-6">
  <div class="md:mt-0 md:col-span-3">
    <div class="px-4 sm:px-0">
      <h3 class="text-lg font-medium leading-6 text-gray-900">
        <slot name="title" />
        {title}
      </h3>
      <p class="mt-1 text-sm leading-5">
        <slot name="description" />
        {description}
      </p>
    </div>
  </div>

  <div
    class="mt-5 md:mt-0 md:col-span-7 shadow sm:rounded-md bg-white {current ? 'border border-mint-500' : ''}">
    {#if header}
      <header
        class="px-4 sm:px-4 py-3 {invalid ? 'bg-red-600' : 'bg-mint-600'} text-gray-50 sm:rounded-t-md">
        <slot name="header" />
      </header>
    {/if}

    <div class="px-4 py-5 sm:p-6">
      <slot />
    </div>

    {#if footer}
      <footer class="px-4 py-3 bg-gray-50 sm:px-6 sm:rounded-b-md">
        <slot name="footer" />
      </footer>
    {/if}
  </div>
</section>
