<script>
  export let title = "";
  export let description = "";
  export let progress = 0;
  export let header = false;
  export let footer = false;
  export let current = false;
</script>

<section class="mt-4 md:grid md:grid-cols-10 md:gap-6">
  <div class="md:mt-0 md:col-span-3">
    <div class="px-4 sm:px-0">
      <h3
        class="mt-3 text-lg leading-6 font-semibold {current ? 'text-mint-800' : 'text-gray-300'}">
        <slot name="title" />
        {title}
      </h3>
      <p
        class="mt-1 text-sm leading-5  {current ? 'text-mint-600' : 'text-gray-200'}">
        <slot name="description" />
        {description}
      </p>
    </div>
  </div>

  <div
    class="mt-5 md:mt-0 md:col-span-7 shadow sm:rounded-lg {current ? 'border-2 border-mint-300 bg-mint-600 text-gray-50' : 'text-mint-600 bg-gray-50'} relative overflow-hidden">
    <div
      class="absolute bg-mint-900 h-full top-0 left-0 z-0"
      style="width: {progress}%" />

    {#if header}
      <header class="relative z-10 px-4 sm:px-4 py-3">
        <slot name="header" />
      </header>
    {/if}

    <slot />

    {#if footer}
      <footer class="px-4 py-3 bg-gray-50 sm:px-6 sm:rounded-b-md">
        <slot name="footer" />
      </footer>
    {/if}
  </div>
</section>
