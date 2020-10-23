<script>
  import ParticipantList from "../components/participants/ParticipantList.svelte";
  import Header from "../layouts/Header.svelte";
  import { GET_ALL_PARTICIPANTS } from "../lib/queries";
</script>

<div class="pt-2 pb-6 md:py-6">
  <Header title="All Participants" on:click />

  <div class="py-4">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
      <ParticipantList type="all" queryArgs={{ query: GET_ALL_PARTICIPANTS }} />
    </div>
  </div>
</div>
