<script>
  import { onDestroy } from "svelte";
  import { createEventDispatcher } from "svelte";

  export let id = null;
  export let value = null;
  export let placeholder = "Select Item";
  export let className = "";
  export let options = [];
  export let multiple = false;
  export let thin = false;

  $: empty = value === undefined || value === null;
  $: valueOption = getValueOption(value);
  $: valueLabel = getLabel(valueOption);

  const dispatch = createEventDispatcher();

  let open = false;

  function getValueOption(val) {
    if (!multiple) {
      return (
        val !== undefined &&
        val !== null &&
        options.find(opt => opt.value === val)
      );
    }

    return (
      val !== undefined &&
      val !== null &&
      val.length > 0 &&
      val.map(v => options.find(opt => opt.value === v))
    );
  }

  function getLabel(val) {
    if (!multiple) {
      return val ? val.label : placeholder;
    }

    return val ? val.map(v => v.label).join(", ") : placeholder;
  }

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

  function isCurrentValue(value, optionValue) {
    if (!multiple) {
      return value === optionValue;
    }
    return value.includes(optionValue) ? true : false;
  }

  onDestroy(() => {
    listenOtherClick(true);
  });
</script>

<div class="{className} relative">
  <span class="inline-block w-full rounded-md {thin ? '' : 'shadow-sm'}">
    <button
      {id}
      on:click={handleOpen}
      type="button"
      aria-haspopup="listbox"
      aria-expanded="true"
      aria-labelledby="listbox-label"
      title={valueOption && valueOption.title}
      class="{thin ? 'rounded pl-2 pr-8' : 'py-2 border rounded-md pl-3 pr-10'}
      cursor-default relative w-full border-gray-300 bg-white text-left
      focus:outline-none focus:shadow-outline-blue focus:border-blue-300
      transition ease-in-out duration-150 sm:text-sm sm:leading-5">
      <span class="block truncate {empty && 'text-gray-400'}">
        {valueLabel}
      </span>
      <span
        class="absolute inset-y-0 right-0 flex items-center pr-0
        pointer-events-none">
        <svg
          class="h-5 w-5 text-gray-400"
          viewBox="0 0 20 20"
          fill="none"
          stroke="currentColor">
          <path
            d="M7 {thin ? '2m0' : '7l3-3 3 3m0'} 6l-3 3-3-3"
            stroke-width="1.5"
            stroke-linecap="round"
            stroke-linejoin="round" />
        </svg>
      </span>
    </button>
  </span>

  {#if open}
    <div class="absolute mt-1 w-full rounded-md bg-white shadow-lg z-50">
      <ul
        tabindex="-1"
        role="listbox"
        aria-labelledby="listbox-label"
        aria-activedescendant="listbox-item-3"
        class="max-h-60 rounded-md py-1 text-base leading-6 shadow-xs
        overflow-auto focus:outline-none sm:text-sm sm:leading-5">

        {#each options as option, i}
          <li
            id="listbox-option-0"
            role="option"
            class="hover:text-white hover:bg-indigo-600 text-gray-900
            cursor-default select-none relative">

            <button
              class="w-full h-full py-2 pl-3 pr-9 focus:outline-none flex
              justify-between"
              title={option.title}
              on:click={() => {
                if (!multiple) {
                  if (value !== option.value) {
                    value = option.value;
                    dispatch('change', { value });
                  }
                } else {
                  if (!value.includes(option.value)) {
                    value.push(option.value);
                  } else {
                    const valueIndex = value.findIndex(v => v === option.value);
                    value.splice(valueIndex, 1);
                  }
                  value = value;
                  dispatch('change', { value });
                }
              }}>

              <span
                class="{isCurrentValue(value, option.value) ? 'font-semibold' : 'font-normal'}
                block truncate">
                {option.label || option.value}
              </span>

              {#if isCurrentValue(value, option.value)}
                <span
                  class="text-indigo-600 absolute inset-y-0 right-0 flex
                  items-center pr-4">
                  <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      fill-rule="evenodd"
                      d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414
                      0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0
                      011.414 0z"
                      clip-rule="evenodd" />
                  </svg>
                </span>
              {/if}

            </button>
          </li>
        {/each}

      </ul>
    </div>
  {/if}
</div>

<style>
  li:hover span {
    color: white;
  }
</style>
