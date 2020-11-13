<script>
  import { cubicIn, cubicOut } from "svelte/easing";
  import { fade, scale } from "svelte/transition";

  export let title = null;
  export let open = false;
  export let color = "green";
  export let button = "Ok";
  export let handleCancel = () => {
    open = false;
  };
  export let handleAccept = () => {
    open = false;
  };
</script>

{#if open}
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
        <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
          {#if title}
            <h3
              class="text-lg leading-6 font-medium text-gray-900"
              id="modal-headline">
              {title}
            </h3>
          {/if}
          <div class="mt-2">
            <slot />
          </div>
        </div>
      </div>
      <div class="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse">
        <span class="flex w-full rounded-md shadow-sm sm:ml-3 sm:w-auto">
          <button
            on:click={handleAccept}
            type="button"
            class="inline-flex justify-center w-full rounded-md border
            border-transparent px-4 py-2 bg-{color}-600 text-base
            leading-6 font-medium text-white shadow-sm hover:bg-{color}-500
            focus:outline-none focus:border-{color}-700
            focus:shadow-outline transition ease-in-out
            duration-150 sm:text-sm sm:leading-5">
            {button}
          </button>
        </span>
        <span class="mt-3 flex w-full rounded-md shadow-sm sm:mt-0 sm:w-auto">
          <button
            on:click={handleCancel}
            type="button"
            class="inline-flex justify-center w-full rounded-md border
            border-gray-300 px-4 py-2 bg-white text-base leading-6 font-medium
            text-gray-700 shadow-sm hover:text-gray-500 focus:outline-none
            focus:border-mint-300 focus:shadow-outline transition
            ease-in-out duration-150 sm:text-sm sm:leading-5">
            Cancel
          </button>
        </span>
      </div>
    </div>
  </div>
{/if}
