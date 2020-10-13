<script>
  import Link from "../base/Link.svelte";
  import Duration from "../misc/Duration.svelte";
  import OptionsMenu from "../misc/OptionsMenu.svelte";
  import RelativeTime from "../misc/RelativeTime.svelte";
  import StatusBadge from "../misc/StatusBadge.svelte";

  export let index = 0;
  export let projectID;
  export let runID;
  export let name;
  // Can be "created", "running", "done", "terminated", "failed"
  export let status = "internal";
  export let startAt = null;
  export let startedAt = null;
  export let endedAt = null;
  export let stepCount = 0;

  const menuOptions = [{ text: "Duplicate", onClick: handleDuplicate }];

  function handleDuplicate() {
    console.log("should dup");
  }
</script>

<li class={index !== 0 && 'border-t border-gray-200'}>
  <Link
    to="/projects/{projectID}/runs/{runID}"
    className="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50
    transition duration-150 ease-in-out sm:rounded-md">
    <div class="flex items-center px-4 py-4 sm:px-2">
      <div class="min-w-0 flex-1 px-4 sm:grid sm:grid-cols-2 sm:gap-4">
        <div>
          <div class="text-sm leading-5 font-medium text-mint-600 truncate">
            {name}
          </div>
          <div
            class="mt-2 mr-6 flex items-center text-sm leading-5 text-gray-500">
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
            Step{stepCount !== 1 ? 's' : ''}
          </div>
        </div>

        <div class="mt-4 sm:mt-0 whitespace-no-wrap">
          <div>
            <div>
              <StatusBadge {status} />
            </div>
            <div class="mt-2 sm:flex">
              {#if endedAt}
                <div
                  class="mr-6 flex items-center text-sm leading-5 text-gray-500">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                    fill="currentColor">
                    <path
                      d="M6.382 5.968A8.962 8.962 0 0 1 12 4c2.125 0 4.078.736
                      5.618 1.968l1.453-1.453 1.414 1.414-1.453 1.453a9 9 0 1
                      1-14.064 0L3.515 5.93l1.414-1.414 1.453 1.453zM12 20a7 7 0
                      1 0 0-14 7 7 0 0 0 0 14zm1-8h3l-5 6.5V14H8l5-6.505V12zM8
                      1h8v2H8V1z" />
                  </svg>
                  <div class="mr-1">
                    {#if status === 'TERMINATED'}
                      Canceled
                    {:else if status === 'FAILED'}Failed{:else}Finished{/if}
                  </div>
                  <RelativeTime time={endedAt} />
                </div>
              {:else if startedAt}
                <div
                  class="mr-6 flex items-center text-sm leading-5 text-gray-500">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                    fill="currentColor">
                    <path
                      d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10
                      10-4.477 10-10 10zm0-2a8 8 0 1 0 0-16 8 8 0 0 0 0
                      16zm1-8h4v2h-6V7h2v5z" />
                  </svg>
                  <div class="mr-1">Started</div>
                  <RelativeTime time={startedAt} />
                </div>
              {:else if startAt}
                <div
                  class="mr-6 flex items-center text-sm leading-5 text-gray-500">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                    fill="currentColor">
                    <path
                      d="M12 2c5.52 0 10 4.48 10 10s-4.48 10-10 10S2 17.52 2 12
                      6.48 2 12 2zm0 18c4.42 0 8-3.58 8-8s-3.58-8-8-8-8 3.58-8 8
                      3.58 8 8 8zm3.536-12.95l1.414 1.414-4.95 4.95L10.586
                      12l4.95-4.95z" />
                  </svg>
                  <div class="mr-1">Starts</div>
                  <RelativeTime time={startAt} />
                </div>
              {:else if status === 'CREATED'}
                <div
                  class="mr-6 flex items-center text-sm leading-5 text-gray-500">
                  Not scheduled
                </div>
              {/if}

              {#if startedAt && endedAt}
                <div
                  class="mr-6 flex items-center text-sm leading-5 text-gray-500">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
                    fill="currentColor">
                    <path
                      d="M17.618 5.968l1.453-1.453 1.414 1.414-1.453 1.453a9 9 0
                      1 1-1.414-1.414zM12 20a7 7 0 1 0 0-14 7 7 0 0 0 0 14zM11
                      8h2v6h-2V8zM8 1h8v2H8V1z" />
                  </svg>
                  <div class="mr-1">Lasted</div>
                  <Duration from={startedAt} to={endedAt} />
                </div>
              {/if}
            </div>
          </div>
        </div>
      </div>
      <div>
        <OptionsMenu className="flex-shrink-0" options={menuOptions} />
      </div>
    </div>
  </Link>
</li>
