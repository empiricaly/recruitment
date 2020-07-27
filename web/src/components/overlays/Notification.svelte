<script context="module">
  import { writable } from "svelte/store";
  import { genID } from "../../utils/uniq.js";

  const notifications = writable([]);

  export function notify(notification) {
    notification.id = genID();
    setTimeout(() => {
      notifications.update(notifs => {
        return notifs.filter(n => n !== notification);
      });
    }, 4000);
    notifications.update(notifs => {
      return [notification, ...notifs];
    });
  }
</script>

<script>
  import { quintOut } from "svelte/easing";
  import { crossfade } from "svelte/transition";
  import { flip } from "svelte/animate";

  const [send, receive] = crossfade({
    duration: d => Math.sqrt(d * 100),

    fallback(node, params) {
      const style = getComputedStyle(node);
      const transform = style.transform === "none" ? "" : style.transform;

      return {
        duration: 300,
        easing: quintOut,
        css: t => `
					transform: ${transform} scale(${t});
					opacity: ${t}
				`
      };
    }
  });

  function handleRemove(notification) {
    return function() {
      notifications.update(notifs => {
        return notifs.filter(n => n !== notification);
      });
    };
  }
</script>

<div
  class="fixed inset-0 flex flex-col items-center justify-start px-4 py-6
  pointer-events-none sm:p-6 sm:items-center sm:justify-start">
  <!--
    Notification panel, show/hide based on alert state.

    Entering: "transform ease-out duration-300 transition"
      From: "translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2"
      To: "translate-y-0 opacity-100 sm:translate-x-0"
    Leaving: "transition ease-in duration-100"
      From: "opacity-100"
      To: "opacity-0"
  -->
  {#each $notifications as notification (notification.id)}
    <div
      in:receive={{ key: notification.id }}
      out:send={{ key: notification.id }}
      animate:flip={{ duration: 200 }}
      class="relative max-w-sm w-full bg-white shadow-lg rounded-lg
      pointer-events-auto mb-4">
      <div class="rounded-lg shadow-xs overflow-hidden">
        <div class="p-4">
          <div class="flex items-start">

            <div class="flex-shrink-0">
              {#if notification.success}
                <svg
                  class="h-6 w-6 text-green-400"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="currentColor"
                  viewBox="0 0 512 512">
                  <path
                    d="M256 8C119.033 8 8 119.033 8 256s111.033 248 248 248
                    248-111.033 248-248S392.967 8 256 8zm0 464c-118.664
                    0-216-96.055-216-216 0-118.663 96.055-216 216-216 118.664 0
                    216 96.055 216 216 0 118.663-96.055 216-216
                    216zm141.63-274.961L217.15 376.071c-4.705 4.667-12.303
                    4.637-16.97-.068l-85.878-86.572c-4.667-4.705-4.637-12.303.068-16.97l8.52-8.451c4.705-4.667
                    12.303-4.637 16.97.068l68.976 69.533
                    163.441-162.13c4.705-4.667 12.303-4.637 16.97.068l8.451
                    8.52c4.668 4.705 4.637 12.303-.068 16.97z" />
                </svg>
              {:else if notification.failed}
                <svg
                  class="h-6 w-6 text-red-400"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="currentColor"
                  viewBox="0 0 512 512">
                  <path
                    d="M256 8C119 8 8 119 8 256s111 248 248 248 248-111
                    248-248S393 8 256 8zm0 464c-118.7 0-216-96.1-216-216 0-118.7
                    96.1-216 216-216 118.7 0 216 96.1 216 216 0 118.7-96.1
                    216-216 216zm94.8-285.3L281.5 256l69.3 69.3c4.7 4.7 4.7 12.3
                    0 17l-8.5 8.5c-4.7 4.7-12.3 4.7-17 0L256 281.5l-69.3
                    69.3c-4.7 4.7-12.3 4.7-17 0l-8.5-8.5c-4.7-4.7-4.7-12.3
                    0-17l69.3-69.3-69.3-69.3c-4.7-4.7-4.7-12.3
                    0-17l8.5-8.5c4.7-4.7 12.3-4.7 17 0l69.3 69.3
                    69.3-69.3c4.7-4.7 12.3-4.7 17 0l8.5 8.5c4.6 4.7 4.6 12.3 0
                    17z" />
                </svg>
              {:else if notification.warning}
                <svg
                  class="h-6 w-6 text-orange-400"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="currentColor"
                  viewBox="0 0 512 512">
                  <path
                    d="M256 40c118.621 0 216 96.075 216 216 0 119.291-96.61
                    216-216 216-119.244 0-216-96.562-216-216 0-119.203
                    96.602-216 216-216m0-32C119.043 8 8 119.083 8 256c0 136.997
                    111.043 248 248 248s248-111.003 248-248C504 119.083 392.957
                    8 256 8zm-11.49 120h22.979c6.823 0 12.274 5.682 11.99
                    12.5l-7 168c-.268 6.428-5.556 11.5-11.99 11.5h-8.979c-6.433
                    0-11.722-5.073-11.99-11.5l-7-168c-.283-6.818 5.167-12.5
                    11.99-12.5zM256 340c-15.464 0-28 12.536-28 28s12.536 28 28
                    28 28-12.536 28-28-12.536-28-28-28z" />
                </svg>
              {:else}
                <svg
                  class="h-6 w-6 text-gray-400"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="currentColor"
                  viewBox="0 0 512 512">
                  <path
                    d="M256 40c118.621 0 216 96.075 216 216 0 119.291-96.61
                    216-216 216-119.244 0-216-96.562-216-216 0-119.203
                    96.602-216 216-216m0-32C119.043 8 8 119.083 8 256c0 136.997
                    111.043 248 248 248s248-111.003 248-248C504 119.083 392.957
                    8 256 8zm-36 344h12V232h-12c-6.627
                    0-12-5.373-12-12v-8c0-6.627 5.373-12 12-12h48c6.627 0 12
                    5.373 12 12v140h12c6.627 0 12 5.373 12 12v8c0 6.627-5.373
                    12-12 12h-72c-6.627 0-12-5.373-12-12v-8c0-6.627 5.373-12
                    12-12zm36-240c-17.673 0-32 14.327-32 32s14.327 32 32 32
                    32-14.327 32-32-14.327-32-32-32z" />
                </svg>
              {/if}
              <!-- <svg
                class="h-6 w-6 text-green-400"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg> -->
            </div>

            <div class="ml-3 w-0 flex-1 pt-0.5">
              <p class="text-sm leading-5 font-medium text-gray-900">
                {notification.title}
              </p>
              <p class="mt-1 text-sm leading-5 text-gray-500">
                {notification.body}
              </p>
            </div>
            <div class="ml-4 flex-shrink-0 flex">
              <button
                on:click={handleRemove(notification)}
                class="inline-flex text-gray-400 focus:outline-none
                focus:text-gray-500 transition ease-in-out duration-150">
                <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                  <path
                    fill-rule="evenodd"
                    d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0
                    111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10
                    11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293
                    5.707a1 1 0 010-1.414z"
                    clip-rule="evenodd" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  {/each}
</div>
