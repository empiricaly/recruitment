<script>
  import { fly } from "svelte/transition";
  import { quintOut } from "svelte/easing";

  export let title = null;
  export let open = false;
</script>

{#if open}
  <div class="fixed inset-0 overflow-hidden" on:click={() => (open = false)}>
    <div class="absolute inset-0 overflow-hidden">
      <section class="absolute inset-y-0 right-0 pl-10 max-w-full flex">
        <!--
        Slide-over panel, show/hide based on slide-over state.

        Entering: "transform transition ease-in-out duration-500 sm:duration-700"
          From: "translate-x-full"
          To: "translate-x-0"
        Leaving: "transform transition ease-in-out duration-500 sm:duration-700"
          From: "translate-x-0"
          To: "translate-x-full"
      -->
        <div
          class="w-screen max-w-md"
          on:click|stopPropagation
          transition:fly={{ duration: 300, x: 500, y: 0, opacity: 1, easing: quintOut }}>
          <div
            class="h-full flex flex-col space-y-6 py-6 bg-white shadow-xl
            overflow-y-scroll">
            <header class="px-4 sm:px-6">
              <div class="flex items-start justify-between space-x-3">
                {#if title}
                  <h2 class="text-lg leading-7 font-medium text-gray-900">
                    {title}
                  </h2>
                {:else}
                  <div />
                {/if}
                <div class="h-7 flex items-center">
                  <button
                    on:click={() => (open = false)}
                    aria-label="Close panel"
                    class="text-gray-400 hover:text-gray-500 transition
                    ease-in-out duration-150">
                    <svg
                      class="h-6 w-6"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor">
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                </div>
              </div>
            </header>
            <div class="relative flex-1 px-4 sm:px-6">
              <slot />
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
{/if}
