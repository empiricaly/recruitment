<script>
  import Layout from "../layouts/Layout.svelte";
  import TemplateLine from "../components/templates/TemplateLine.svelte";
</script>

<Layout title="Templates" action="New Template">
  <div class="bg-white shadow sm:rounded-md">
    <ul>
      <TemplateLine
        name="Speed Dating Pilot 2"
        type="internal"
        path="/"
        stepCount={2}
        runCount={0}
        creator="Nicolas" />
      <TemplateLine
        index="1"
        name="Speed Dating Pilot 2"
        type="internal"
        path="/"
        stepCount={2}
        runCount={0}
        creator="Nicolas" />
      <TemplateLine
        index="2"
        name="Speed Dating Pilot 1"
        type="mturk"
        path="/"
        stepCount={4}
        runCount={5}
        creator="Nicolas" />
    </ul>
  </div>
</Layout>
