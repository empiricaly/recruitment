<script>
  import Layout from "../layouts/Layout.svelte";
  import TemplateLine from "../components/templates/TemplateLine.svelte";
</script>

<Layout title="Templates" action="New Template">
  <div class="bg-white shadow sm:rounded-md">
    <ul>
      <TemplateLine
        projectID="1234567"
        templateID="2345678"
        name="Speed Dating Pilot 2"
        type="internal"
        stepCount={2}
        runCount={0}
        creator="Nicolas" />
      <TemplateLine
        projectID="1234567"
        templateID="2345678"
        index="1"
        name="Speed Dating Pilot 2"
        type="internal"
        stepCount={2}
        runCount={0}
        creator="Nicolas" />
      <TemplateLine
        projectID="1234567"
        templateID="2345678"
        index="2"
        name="Speed Dating Pilot 1"
        type="mturk"
        stepCount={4}
        runCount={5}
        creator="Nicolas" />
    </ul>
  </div>
</Layout>
