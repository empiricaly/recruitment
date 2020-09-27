<script>
  import Template from "../components/templates/Template.svelte";
  import Layout from "../layouts/Layout.svelte";

  let template = {
    projectID: "fkelbwkjq",
    name: "Speed Dating 1",
    selectionType: "INTERNAL_DB",
    participantCount: null,
    adult: false,
    mturkCriteria: { qualifications: [] },
    internalCriteria: {
      condition: {
        and: [
          {
            key: "",
            comparator: "EQUAL_TO",
            values: [],
          },
        ],
      },
    },
    steps: [],
  };

  // $: console.log(JSON.stringify(template, "", "  "));
  $: console.log(JSON.stringify(template));
</script>

<Layout bind:title={template.name} overtitle="Template" titleUpdatable>
  <Template bind:template />
</Layout>
