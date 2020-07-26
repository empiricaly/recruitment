<script context="module">
  import { get, writable } from "svelte/store";

  const open = writable(false);
  const config = writable({});
  const prom = writable(null);

  const defaultConfig = {
    color: "red",
    title: null,
    button: "Ok"
  };

  export function confirm(conf) {
    if (get(prom)) {
      throw "Alert already open";
    }
    if (!conf.body) {
      throw "body required for Alert";
    }

    config.set({ ...defaultConfig, ...conf });
    open.set(true);

    return new Promise((resolve, reject) => {
      prom.set({ resolve, reject });
    });
  }
</script>

<script>
  import { fade, scale } from "svelte/transition";
  import { cubicIn, cubicOut } from "svelte/easing";

  function handleCancel() {
    $open = false;
    $prom.reject();
    $prom = null;
    config.set(defaultConfig);
  }

  function handleAccept() {
    $open = false;
    $prom.resolve();
    $prom = null;
    config.set(defaultConfig);
  }
</script>

{#if $open}
  <div
    class="fixed bottom-0 inset-x-0 px-4 pb-4 sm:inset-0 sm:flex sm:items-center
    sm:justify-center">
    <div
      class="fixed inset-0 transition-opacity"
      in:fade={{ duration: 300, easing: cubicOut }}
      out:fade={{ duration: 200, easing: cubicIn }}
      on:click={handleCancel}>
      <div class="absolute inset-0 bg-gray-500 opacity-75" />
    </div>
    <div
      in:scale={{ start: 0.95, duration: 300, easing: cubicOut }}
      out:scale={{ start: 0.95, duration: 200, easing: cubicIn }}
      class="bg-white rounded-lg px-4 pt-5 pb-4 overflow-hidden shadow-xl
      transform transition-all sm:max-w-lg sm:w-full sm:p-6 "
      role="dialog"
      aria-modal="true"
      aria-labelledby="modal-headline">
      <div class="sm:flex sm:items-start">
        <div
          class="mx-auto flex-shrink-0 flex items-center justify-center h-12
          w-12 rounded-full bg-{$config.color}-100 sm:mx-0 sm:h-10 sm:w-10">
          <svg
            class="h-6 w-6 text-{$config.color}-600"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667
              1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77
              1.333.192 3 1.732 3z" />
          </svg>
        </div>
        <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
          {#if $config.title}
            <h3
              class="text-lg leading-6 font-medium text-gray-900"
              id="modal-headline">
              {$config.title}
            </h3>
          {/if}
          <div class="mt-2">
            <p class="text-sm leading-5 text-gray-500">{$config.body}</p>
          </div>
        </div>
      </div>
      <div class="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
        <span class="flex w-full rounded-md shadow-sm sm:ml-3 sm:w-auto">
          <button
            on:click={handleAccept}
            type="button"
            class="inline-flex justify-center w-full rounded-md border
            border-transparent px-4 py-2 bg-{$config.color}-600 text-base
            leading-6 font-medium text-white shadow-sm hover:bg-{$config.color}-500
            focus:outline-none focus:border-{$config.color}-700
            focus:shadow-outline-{$config.color} transition ease-in-out
            duration-150 sm:text-sm sm:leading-5">
            {$config.button}
          </button>
        </span>
        <span class="mt-3 flex w-full rounded-md shadow-sm sm:mt-0 sm:w-auto">
          <button
            on:click={handleCancel}
            type="button"
            class="inline-flex justify-center w-full rounded-md border
            border-gray-300 px-4 py-2 bg-white text-base leading-6 font-medium
            text-gray-700 shadow-sm hover:text-gray-500 focus:outline-none
            focus:border-blue-300 focus:shadow-outline-blue transition
            ease-in-out duration-150 sm:text-sm sm:leading-5">
            Cancel
          </button>
        </span>
      </div>
    </div>
  </div>
{/if}
