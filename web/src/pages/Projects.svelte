<script context="module">
  function camelize(str) {
    return str
      .replace(/(?:^\w|[A-Z]|\b\w)/g, function (word, index) {
        return index === 0 ? word.toLowerCase() : word.toUpperCase();
      })
      .replace(/\W/g, "")
      .replace(/\s+/g, "");
  }
</script>

<script>
  import { mutate, query } from "svelte-apollo";
  import Button from "../components/base/Button.svelte";
  import Callout from "../components/base/Callout.svelte";
  import Input from "../components/base/Input.svelte";
  import Label from "../components/base/Label.svelte";
  import LinkButton from "../components/base/LinkButton.svelte";
  import { notify } from "../components/overlays/Notification.svelte";
  import ProjectLine from "../components/projects/ProjectLine.svelte";
  import { client } from "../lib/apollo";
  import { CREATE_PROJECT, GET_PROJECTS } from "../lib/queries";
  import { handleErrorMessage } from "../utils/errorQuery";

  let newProject = false;
  let name = "";
  let projectID = "";
  let idGotFocused = false;

  $: {
    if (!idGotFocused) {
      projectID = camelize(name);
    }
  }

  const camelCasePattern = "([a-z]+[A-Z]*\\w+)+";

  $: projects = query(client, { query: GET_PROJECTS });

  async function handleNewProject(event) {
    event.preventDefault();

    try {
      await mutate(client, {
        mutation: CREATE_PROJECT,
        variables: {
          input: {
            projectID,
            name,
          },
        },
      });
      projects.refetch();
      newProject = false;
      name = "";
      projectID = "";
    } catch (error) {
      handleErrorMessage(error);
      notify({
        failed: true,
        title: `Could not create Project`,
        body: error.graphQLErrors.map((e) => e.message).join(", "),
      });
    }
  }
</script>

<main class="flex justify-center items-center h-full w-full">
  {#await $projects}
    Loading...
  {:then result}
    {#if result.data.projects.length === 0}
      <div class="w-64 px-4 py-4">
        <Callout color="yellow">
          You have no projects yet.
          <br />
          Create one now!
        </Callout>
      </div>
    {/if}
    <div class="h-full">
      {#if newProject || result.data.projects.length === 0}
        <div class="h-full flex items-center">
          <div class="w-128 px-4 py-4">
            <h1 class="font-semibold">New Project</h1>
            <form class="mt-4" on:submit={handleNewProject}>
              <section>
                <Label forID="name" text="Project Name" />
                <Input
                  focus
                  bind:value={name}
                  id="name"
                  required
                  placeholder="Enter Project Name" />
                <div class="mt-2 text-sm text-gray-400 leading-tight">
                  The Project Name is the human-friendly name for your Project.
                </div>
              </section>
              <section class="mt-4">
                <Label forID="projectID" text="Project Identifier" />
                <Input
                  on:focus={() => (idGotFocused = true)}
                  bind:value={projectID}
                  id="projectID"
                  pattern={camelCasePattern}
                  required
                  placeholder="Enter Project Identifier" />
                <div class="mt-2 text-sm text-gray-400 leading-tight">
                  The Project Identifier is used to label Participants having
                  taken part in the Project. It should be written in Camel Case,
                  e.g., myCoolProject.
                </div>
              </section>
              <div class="mt-6 w-64 flex">
                <div>
                  <Button type="submit" text="Create Project" />
                </div>

                {#if result.data.projects.length !== 0}
                  <div class="ml-3">
                    <Button
                      on:click={() => (newProject = false)}
                      text="Cancel"
                      secondary />
                  </div>
                {/if}
              </div>
            </form>
          </div>
        </div>
      {:else}
        <div
          class="flex flex-col sm:flex-row justify-center items-stretch sm:items-center h-full overflow-hidden">
          <div class="max-h-full overflow-auto">
            <div class="sm:w-64 md:w-96 px-8 py-4">
              <h1 class="font-semibold sm:pl-2">Projects</h1>
              <ul class="mt-3">
                {#each result.data.projects as project}
                  <ProjectLine id={project.projectID} name={project.name} />
                {/each}
              </ul>

              <div class="mt-8">
                <Button
                  on:click={() => (newProject = true)}
                  text="Create a new Project"
                  secondary
                  full />
              </div>
            </div>
          </div>

          <div class="mt-8 sm:mt-0 sm:w-64 md:w-96 px-8 py-4">
            <h1 class="font-semibold sm:pl-2">Participants</h1>
            <div class="mt-3">See all participants in the internal DB.</div>
            <div class="mt-8">
              <LinkButton to="/participants" secondary>
                View Participants
              </LinkButton>
            </div>
          </div>
        </div>
      {/if}
    </div>
  {:catch error}
    Error loading Projects:
    {error}
    {handleErrorMessage(error)}
  {/await}
</main>
