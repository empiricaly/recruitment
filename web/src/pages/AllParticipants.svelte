<script>
  import Header from "../layouts/Header.svelte";
  import ParticipantList from "../components/participants/ParticipantList.svelte";
</script>

<div class="pt-2 pb-6 md:py-6">
  <Header title="All Participants" on:click />

  <div class="py-4">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 md:px-8">
      <ParticipantList />
    </div>
  </div>
</div>
