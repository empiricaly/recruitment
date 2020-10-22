<script>
  import { query } from "svelte-apollo";
  import { client } from "../../lib/apollo";
  import { GET_PARTICIPANTS } from "../../lib/queries";
  import DataCell from "./DataCell.svelte";

  export let project;
  export let run;

  console.log(run);

  const participantsQuery = query(client, {
    query: GET_PARTICIPANTS,
    variables: { projectID: project.projectID, runID: run.id },
  });

  let participants;
  $: $participantsQuery.then((result) => {
    const steps = result.data.project.runs[0].steps;
    if (steps.length === 0) {
      return;
    }
    participants = steps[0].participants;
  });

  let keys = [];
  $: if (participants && participants.length > 0) {
    const lkeys = {};
    for (let i = 0; i < participants.length; i++) {
      for (let j = 0; j < participants[i].data.length; j++) {
        lkeys[participants[i].data[j].key] = true;
      }
    }
    keys = Object.keys(lkeys);
  }

  $: console.log(participants);
  $: console.log(keys);

  $: console.log(participants);
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
                    class="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider">
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
        </div>
      </div>
    </div>
  </div>
{:else}Loading...{/if}
