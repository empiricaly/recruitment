<script>
  import { createEventDispatcher } from "svelte";
  import Button from "../components/base/Button.svelte";
  import OptionsMenu from "../components/misc/OptionsMenu.svelte";

  const dispatch = createEventDispatcher();

  export let project;
  export let title = "";
  export let titleUpdatable = false;
  export let overtitle = "";
  export let action = "";
  export let facts = [];
  export let actions = [];

  $: secondaryActions = actions.filter((a) => !a.primary);
  let prevTitle = title;

  function handleTitleChange() {
    if (!title) {
      title = prevTitle;
      return;
    }

    prevTitle = title;
  }

  // export let facts = [
  //   {
  //     text: `Full-time`,
  //     icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M464 128H352V56c0-13.26-10.74-24-24-24H184c-13.26 0-24 10.74-24 24v72H48c-26.51 0-48 21.49-48 48v256c0 26.51 21.49 48 48 48h416c26.51 0 48-21.49 48-48V176c0-26.51-21.49-48-48-48zM192 64h128v64H192V64zm288 368c0 8.82-7.18 16-16 16H48c-8.82 0-16-7.18-16-16V288h160v40c0 13.25 10.75 24 24 24h80c13.25 0 24-10.75 24-24v-40h160v144zM224 320v-32h64v32h-64zm256-64H32v-80c0-8.82 7.18-16 16-16h416c8.82 0 16 7.18 16 16v80z"/></svg>`
  //   },
  //   {
  //     text: `Remote`,
  //     icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512"><path d="M192 0C85.961 0 0 85.961 0 192c0 77.413 26.97 99.031 172.268 309.67 9.534 13.772 29.929 13.774 39.465 0C357.03 291.031 384 269.413 384 192 384 85.961 298.039 0 192 0zm0 473.931C52.705 272.488 32 256.494 32 192c0-42.738 16.643-82.917 46.863-113.137S149.262 32 192 32s82.917 16.643 113.137 46.863S352 149.262 352 192c0 64.49-20.692 80.47-160 281.931z"/></svg>`
  //   },
  //   {
  //     text: `$120k – $140k`,
  //     icon: `<path
  //           d="M8.433 7.418c.155-.103.346-.196.567-.267v1.698a2.305 2.305 0
  //           01-.567-.267C8.07 8.34 8 8.114 8 8c0-.114.07-.34.433-.582zM11
  //           12.849v-1.698c.22.071.412.164.567.267.364.243.433.468.433.582 0
  //           .114-.07.34-.433.582a2.305 2.305 0 01-.567.267z" />
  //         <path
  //           fill-rule="evenodd"
  //           d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-13a1 1 0 10-2 0v.092a4.535
  //           4.535 0 00-1.676.662C6.602 6.234 6 7.009 6 8c0 .99.602 1.765 1.324
  //           2.246.48.32 1.054.545
  //           1.676.662v1.941c-.391-.127-.68-.317-.843-.504a1 1 0 10-1.51
  //           1.31c.562.649 1.413 1.076 2.353 1.253V15a1 1 0 102 0v-.092a4.535
  //           4.535 0 001.676-.662C13.398 13.766 14 12.991 14
  //           12c0-.99-.602-1.765-1.324-2.246A4.535 4.535 0 0011
  //           9.092V7.151c.391.127.68.317.843.504a1 1 0
  //           101.511-1.31c-.563-.649-1.413-1.076-2.354-1.253V5z"
  //           clip-rule="evenodd" />`
  //   }
  // ];
  // export let actions = [
  //   {
  //     text: "Edit",
  //     action: "edit",
  //     icon: `<path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />`
  //   },
  //   {
  //     text: "View",
  //     action: "view",
  //     icon: `<path fill-rule="evenodd" d="M12.586 4.586a2 2 0 112.828 2.828l-3 3a2 2 0 01-2.828 0 1 1 0 00-1.414 1.414 4 4 0 005.656 0l3-3a4 4 0 00-5.656-5.656l-1.5 1.5a1 1 0 101.414 1.414l1.5-1.5zm-5 5a2 2 0 012.828 0 1 1 0 101.414-1.414 4 4 0 00-5.656 0l-3 3a4 4 0 105.656 5.656l1.5-1.5a1 1 0 10-1.414-1.414l-1.5 1.5a2 2 0 11-2.828-2.828l3-3z" clip-rule="evenodd" />`
  //   },
  //   {
  //     text: "Publish",
  //     action: "publish",
  //     icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M435.848 83.466L172.804 346.51l-96.652-96.652c-4.686-4.686-12.284-4.686-16.971 0l-28.284 28.284c-4.686 4.686-4.686 12.284 0 16.971l133.421 133.421c4.686 4.686 12.284 4.686 16.971 0l299.813-299.813c4.686-4.686 4.686-12.284 0-16.971l-28.284-28.284c-4.686-4.686-12.284-4.686-16.97 0z"/></svg>`,
  //     primary: true
  //   }
  // ];
</script>

<div
  class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 lg:flex lg:items-start
    lg:justify-between">
  {#if overtitle}
    <div class="uppercase tracking-wide font-medium text-gray-500 mr-1 text-sm">
      {overtitle}
    </div>
  {/if}
  <slot name="pretitle" />
</div>

<div
  class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 lg:flex lg:items-baseline
    lg:justify-between">
  <div class="flex-1 min-w-0">
    <h1
      class="mt-1 text-2xl font-bold leading-7 text-gray-900 sm:text-3xl
        sm:leading-9 sm:truncate flex items-center">
      {#if titleUpdatable}
        <input
          class="font-bold focus:outline-none
            rounded-md bg-transparent focus:bg-white w-full"
          type="text"
          bind:value={title}
          on:change={handleTitleChange} />
      {:else}
        <div>{title}</div>
      {/if}
      <slot name="posttitle" />
    </h1>
    <div class="mt-1 flex flex-col sm:mt-0 sm:flex-row sm:flex-wrap">
      {#each facts as fact, i (fact.text)}
        <div
          class="mt-2 flex items-center text-sm leading-5 text-gray-500 sm:mr-6">
          <svg
            class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
            viewBox="0 0 20 20"
            fill="currentColor">
            {@html fact.icon}
          </svg>
          {fact.text}
        </div>
      {/each}
    </div>
  </div>

  {#if action}
    <Button
      on:click={() => dispatch('click', { action: 'main', project })}
      text={action} />
  {/if}
  <slot name="actions" />

  {#if actions.length > 0}
    <div class="mt-5 flex lg:mt-0 lg:ml-4">
      {#each actions as action, index (action.text)}
        <span
          class="{action.primary ? '' : 'hidden sm:block'}
            {index > 0 ? 'ml-3' : ''} shadow-sm rounded-md">
          <Button
            on:click={() => dispatch('click', {
                action: action.action,
                project,
              })}
            disabled={action.disabled}
            tertiary={!action.primary}
            icon={action.icon}
            text={action.text} />
        </span>
      {/each}

      <span class="ml-3 relative shadow-sm rounded-md sm:hidden">
        <OptionsMenu
          on:click
          className="flex-shrink-0"
          options={secondaryActions}
          let:handleOpen>
          <Button
            tertiary
            text="More"
            on:click={handleOpen}
            icon={`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"></Button><path d="M441.9 167.3l-19.8-19.8c-4.7-4.7-12.3-4.7-17 0L224 328.2 42.9 147.5c-4.7-4.7-12.3-4.7-17 0L6.1 167.3c-4.7 4.7-4.7 12.3 0 17l209.4 209.4c4.7 4.7 12.3 4.7 17 0l209.4-209.4c4.7-4.7 4.7-12.3 0-17z"/></svg>`} />
        </OptionsMenu>
      </span>
    </div>
  {/if}
</div>
