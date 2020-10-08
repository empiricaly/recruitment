<script>
  import Badge from "../base/Badge.svelte";
  import Link from "../base/Link.svelte";
  import CreatedBy from "../misc/CreatedBy.svelte";
  import OptionsMenu from "../misc/OptionsMenu.svelte";

  export let index = 0;
  export let projectID;
  export let templateID;
  export let name;
  export let type = "internal";

  export let stepCount = 0;
  export let runCount = 0;
  export let creator = null;

  const menuOptions = [{ text: "Duplicate", onClick: handleDuplicate }];

  function handleDuplicate() {
    console.log("should dup");
  }
</script>

<li class={index !== 0 && 'border-t border-gray-200'}>
  <Link
    to="/projects/{projectID}/templates/{templateID}"
    className="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50
    transition duration-150 ease-in-out sm:rounded-md">
    <div class="flex items-strect">
      <div class="pl-4 py-4 sm:pl-6 flex-1">
        <div class="flex items-center justify-between">
          <div class="text-sm leading-5 font-medium text-mint-600 truncate">
            {name}
          </div>
          <div class="ml-2 flex-shrink-0 flex">
            {#if type === 'internal'}
              <Badge text="Internal DB" color="teal" />
            {:else if type === 'mturk'}
              <Badge text="AWS Qualifications" color="yellow" />
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
              {stepCount}
              Step{stepCount !== 1 && 's'}
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
              {runCount}
              Run{runCount !== 1 && 's'}
            </div>
          </div>
          <CreatedBy {creator} />
        </div>
      </div>
      <OptionsMenu className="flex-shrink-0" options={menuOptions} />
    </div>
  </Link>
</li>
