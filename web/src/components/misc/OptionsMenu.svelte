<script>
  import { createEventDispatcher,onDestroy } from "svelte";
import { push } from "../../lib/routing.js";


  const dispatch = createEventDispatcher();

  export let className = "";
  export let label = "Options";
  export let options = [];

  let open = false;

  function handleOpen(event) {
    event.stopPropagation();
    event.preventDefault();
    open = !open;
    listenOtherClick();
  }

  function close() {
    open = false;
  }

  function listenOtherClick(mustClose = false) {
    if (open && !mustClose) {
      document.body.addEventListener("click", close);
    } else {
      document.body.removeEventListener("click", close);
    }
  }

  onDestroy(() => {
    listenOtherClick(true);
  });
</script>

<div class="{className} min-h-full">
  <div class="relative h-full">
    <slot {handleOpen}>
      <button
        on:click={handleOpen}
        aria-label={label}
        aria-haspopup="true"
        aria-expanded={open}
        class="flex px-3 h-full items-center text-gray-400 hover:text-gray-600
          focus:outline-none focus:text-gray-600 z-0">
        <svg class="h-5 w-5" fill="currentColor" viewBox="0 0 20 20">
          <path
            d="M10 6a2 2 0 110-4 2 2 0 010 4zM10 12a2 2 0 110-4 2 2 0 010 4zM10
            18a2 2 0 110-4 2 2 0 010 4z" />
        </svg>
      </button>
    </slot>

    {#if open}
      <div
        class="origin-top-right absolute top-0 right-0 mt-12 w-56 rounded-md
          shadow-lg z-10">
        <div class="rounded-md bg-white shadow-xs">
          <div class="py-1" role="menu" aria-orientation="vertical">
            {#each options as option, i}
              <button
                on:click={(event) => {
                  open = false;
                  event.preventDefault();
                  event.stopPropagation();
                  if (option.action) {
                    dispatch('click', { action: option.action });
                  } else if (option.onClick) {
                    option.onClick(event);
                  } else if (option.path) {
                    push(option.path);
                  }
                }}
                href="#"
                class="block px-4 py-2 text-sm leading-5 text-gray-700
                  hover:bg-gray-100 hover:text-gray-900 focus:outline-none
                  focus:bg-gray-100 focus:text-gray-900 w-full"
                role="menuitem">
                {option.text}
              </button>
            {/each}
          </div>
        </div>
      </div>
    {/if}
  </div>
</div>
