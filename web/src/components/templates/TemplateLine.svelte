<script>
  import Link from "../base/Link.svelte";
  import OptionsMenu from "../utils/OptionsMenu.svelte";

  export let index = 0;
  export let name;
  export let type = "internal";
  export let path = "/";
  export let stepCount = 0;
  export let runCount = 0;
  export let creator = null;

  const menuOptions = [{ title: "Duplicate", onClick: handleDuplicate }];

  function handleDuplicate() {
    console.log("should dup");
  }
</script>

<li class={index !== 0 && 'border-t border-gray-200'}>
  <Link
    to={path}
    className="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50
    transition duration-150 ease-in-out sm:rounded-md">
    <div class="flex items-strect">
      <div class="pl-4 py-4 sm:pl-6 flex-1">
        <div class="flex items-center justify-between">
          <div class="text-sm leading-5 font-medium text-indigo-600 truncate">
            {name}
          </div>
          <div class="ml-2 flex-shrink-0 flex">
            {#if type === 'internal'}
              <span
                class="px-2 inline-flex text-xs leading-5 font-semibold
                rounded-full bg-teal-100 text-teal-800">
                Internal DB
              </span>
            {:else if type === 'mturk'}
              <span
                class="px-2 inline-flex text-xs leading-5 font-semibold
                rounded-full bg-yellow-100 text-yellow-800">
                AWS Qualifications
              </span>
            {/if}
          </div>
        </div>
        <div class="mt-2 sm:flex sm:justify-between">
          <div class="sm:flex">
            <div class="mr-6 flex items-center text-sm leading-5 text-gray-500">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                fill="currentColor">
                <path
                  d="M11 4h10v2H11V4zm0 4h6v2h-6V8zm0 6h10v2H11v-2zm0
                  4h6v2h-6v-2zM3 4h6v6H3V4zm2 2v2h2V6H5zm-2 8h6v6H3v-6zm2
                  2v2h2v-2H5z" />
              </svg>
              {stepCount} Step{stepCount !== 1 && 's'}
            </div>
            <div class="mr-6 flex items-center text-sm leading-5 text-gray-500">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                fill="currentColor">
                <path
                  d="M16.394 12L10 7.737v8.526L16.394 12zm2.982.416L8.777
                  19.482A.5.5 0 0 1 8 19.066V4.934a.5.5 0 0 1 .777-.416l10.599
                  7.066a.5.5 0 0 1 0 .832z" />
              </svg>
              {runCount} Run{runCount !== 1 && 's'}
            </div>
          </div>
          <div
            class="mt-2 flex items-center text-sm leading-5 text-gray-500
            sm:mt-0">
            <svg
              class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
              viewBox="0 0 24 24"
              fill="currentColor">
              <path
                d="M12 14v2a6 6 0 0 0-6 6H4a8 8 0 0 1 8-8zm0-1c-3.315
                0-6-2.685-6-6s2.685-6 6-6 6 2.685 6 6-2.685 6-6 6zm0-2c2.21 0
                4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm2.595 7.812a3.51
                3.51 0 0 1 0-1.623l-.992-.573 1-1.732.992.573A3.496 3.496 0 0 1
                17 14.645V13.5h2v1.145c.532.158 1.012.44 1.405.812l.992-.573 1
                1.732-.992.573a3.51 3.51 0 0 1 0 1.622l.992.573-1
                1.732-.992-.573a3.496 3.496 0 0 1-1.405.812V22.5h-2v-1.145a3.496
                3.496 0 0 1-1.405-.812l-.992.573-1-1.732.992-.572zM18 19.5a1.5
                1.5 0 1 0 0-3 1.5 1.5 0 0 0 0 3z" />
            </svg>
            <span>
              {#if creator}Created by {creator}{/if}
              <time datetime="2020-01-07">today</time>
            </span>
          </div>
        </div>
      </div>
      <OptionsMenu className="flex-shrink-0" options={menuOptions} />
    </div>
  </Link>
</li>
