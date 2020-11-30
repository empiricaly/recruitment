<script>
  import { createEventDispatcher } from "svelte";

  export let currentPage = 0;
  export let total = 0;
  export let limit = 0;

  const dispatch = createEventDispatcher();
  let totalPage = Math.ceil(total / limit);
  let pages = [];

  function handleChangePage(nextPage) {
    if (nextPage === "...") {
      return;
    }

    dispatch("changePage", {
      nextPage,
    });
  }

  $: {
    pages = [];

    if (totalPage < 6) {
      for (let i = 0; i < totalPage; i++) {
        pages.push({ number: i + 1, active: i === currentPage });
      }
    } else {
      const rightPage = totalPage - 3;
      const overRightPage = currentPage + 2 < rightPage;
      const leftPage = !overRightPage
        ? { start: rightPage - 3, end: rightPage - 1 }
        : {
            start: currentPage - 2 < 0 ? 0 : currentPage,
            end: currentPage - 2 < 0 ? 2 : currentPage + 2,
          };

      for (let i = leftPage.start; i <= leftPage.end; i++) {
        pages.push({ number: i + 1, active: i === currentPage });
      }

      if (currentPage + 2 < rightPage - 1) {
        pages.push({ number: "...", active: false });
      }

      for (let i = rightPage; i < totalPage; i++) {
        pages.push({ number: i + 1, active: i === currentPage });
      }
    }
  }
</script>

<div
  class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6">
  <div class="flex-1 flex justify-between sm:hidden">
    <button
      disabled={currentPage === 0}
      on:click={() => handleChangePage(currentPage - 1)}
      class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm leading-5 font-medium rounded-md text-gray-700 bg-white hover:text-gray-500 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 active:bg-gray-100 active:text-gray-700 transition ease-in-out duration-150">
      Previous
    </button>
    <button
      disabled={currentPage === totalPage - 1}
      on:click={() => handleChangePage(currentPage + 1)}
      class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm leading-5 font-medium rounded-md text-gray-700 bg-white hover:text-gray-500 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 active:bg-gray-100 active:text-gray-700 transition ease-in-out duration-150">
      Next
    </button>
  </div>
  <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
    <div>
      <p class="text-sm leading-5 text-gray-700">
        Showing
        <span class="font-medium">{currentPage * limit + 1}</span>
        to
        <span
          class="font-medium">{currentPage !== totalPage - 1 ? limit * (currentPage + 1) : total}</span>
        of
        <span class="font-medium">{total}</span>
        results
      </p>
    </div>
    <div>
      <nav class="relative z-0 inline-flex shadow-sm">
        <button
          disabled={currentPage === 0}
          on:click={() => handleChangePage(currentPage - 1)}
          class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm leading-5 font-medium text-gray-500 hover:text-gray-400 focus:z-10 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:bg-gray-100 active:text-gray-500 transition ease-in-out duration-150"
          aria-label="Previous">
          <!-- Heroicon name: chevron-left -->
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor">
            <path
              fill-rule="evenodd"
              d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
              clip-rule="evenodd" />
          </svg>
        </button>
        {#each pages as page}
          <button
            disabled={currentPage + 1 === page.number}
            on:click={() => handleChangePage(page.number - 1)}
            class="{page.active ? 'text-mint-400' : 'text-gray-800'} -ml-px relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm leading-5 font-medium hover:text-gray-500 focus:z-10 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:bg-gray-100 active:text-gray-700 transition ease-in-out duration-150">
            {page.number}
          </button>
        {/each}
        <button
          disabled={currentPage === totalPage - 1}
          on:click={() => handleChangePage(currentPage + 1)}
          class="-ml-px relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm leading-5 font-medium text-gray-500 hover:text-gray-400 focus:z-10 focus:outline-none focus:border-blue-300 focus:shadow-outline-blue active:bg-gray-100 active:text-gray-500 transition ease-in-out duration-150"
          aria-label="Next">
          <!-- Heroicon name: chevron-right -->
          <svg
            class="h-5 w-5"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor">
            <path
              fill-rule="evenodd"
              d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
              clip-rule="evenodd" />
          </svg>
        </button>
      </nav>
    </div>
  </div>
</div>
