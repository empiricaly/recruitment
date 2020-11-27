<script>
  import { query } from "svelte-apollo";
  import { client } from "../../lib/apollo";
  import { participantPerQueryType } from "../../lib/models/participants/participants.js";
  import { handleErrorMessage } from "../../utils/errorQuery";
  import Pagination from "../base/Pagination.svelte";
  import DataCell from "./DataCell.svelte";

  export let queryArgs;
  export let type = "run";
  export let limit = 20;
  export let participants;
  export let keys = [];

  let offset = 0;
  let total = 0;

  $: queryArgs.variables = { ...queryArgs.variables, offset, limit };

  $: participantsQuery = query(client, queryArgs);

  $: try {
    $participantsQuery.then((result) => {
      const pp = participantPerQueryType(type, result);
      if (pp) {
        participants = pp.participants;
        total = pp.total;
      }
    });
  } catch (error) {
    handleErrorMessage(error);
  }

  $: if (participants && participants.length > 0) {
    const lkeys = {};
    for (let i = 0; i < participants.length; i++) {
      for (let j = 0; j < participants[i].data.length; j++) {
        lkeys[participants[i].data[j].key] = true;
      }
    }
    keys = Object.keys(lkeys);
  }

  function handleChangePage(event) {
    const { nextPage } = event.detail;
    offset = nextPage;
  }
</script>

{#if participants}
  <div class="flex flex-col">
    <div class="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
      <div class="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
        <div
          class="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
          <table class="min-w-full divide-y divide-gray-200">
            <thead>
              <tr>
                <th
                  class="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider">
                  WorkerID
                </th>
                {#each keys as key (key)}
                  <th
                    class="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500">
                    {key}
                  </th>
                {/each}
                <!-- <th class="px-6 py-3 bg-gray-50" /> -->
              </tr>
            </thead>
            <tbody>
              {#each participants as participant, i (participant.id)}
                <tr class={i % 2 == 1 ? 'bg-white' : 'bg-gray-50'}>
                  <td
                    class="px-6 py-4 whitespace-no-wrap text-sm leading-5 font-medium text-gray-900">
                    {participant.mturkWorkerID}
                  </td>
                  {#each keys as key (key)}
                    <td
                      class="px-6 py-4 whitespace-no-wrap text-sm leading-5 text-gray-500">
                      <DataCell {participant} {key} />
                    </td>
                  {/each}
                  <!-- <td
                  class="px-6 py-4 whitespace-no-wrap text-right text-sm leading-5 font-medium">
                  <a
                    href="#"
                    class="text-indigo-600 hover:text-indigo-900">Edit</a>
                </td> -->
                </tr>
              {/each}
              <!-- More rows... -->
            </tbody>
          </table>
          <Pagination
            on:changePage={handleChangePage}
            currentPage={offset}
            {limit}
            {total} />
        </div>
      </div>
    </div>
  </div>
{:else}Loading...{/if}
