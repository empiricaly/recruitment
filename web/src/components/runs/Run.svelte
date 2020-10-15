<script>
  import dayjs from "dayjs";
  import { mutate } from "svelte-apollo";
  import Layout from "../../layouts/Layout.svelte";
  import { client } from "../../lib/apollo";
  import { START_RUN, UPDATE_RUN } from "../../lib/queries";
  import { deepCopy } from "../../utils/copy";
  import { debounce } from "../../utils/timing";
  import StatusBadge from "../misc/StatusBadge.svelte";
  import { notify } from "../overlays/Notification.svelte";
  import Template from "../templates/Template.svelte";
  import RunningRun from "./RunningRun.svelte";

  export let project;
  export let run;

  let name = run.name;
  let initialName = name;

  $: {
    if (name !== initialName) {
      update();
      initialName = name;
      console.log("wtf");
    }
  }

  const update = debounce(
    async () => {
      console.log("project", project);
      console.log("run", run);
      try {
        const input = {
          ID: run.id,
          projectID: project.id,
          name,
        };

        console.log(JSON.stringify(input, null, "  "));

        await mutate(client, {
          mutation: UPDATE_RUN,
          variables: {
            input,
          },
        });

        notify({
          success: true,
          title: `Run Saved`,
          // body:
          //   "Something happened on the server, and we could not create a new Run as requested.",
        });
      } catch (error) {
        console.error(error);
        notify({
          failed: true,
          title: `Could not save Run update`,
          body:
            "Something happened on the server, and we could not save the latest changes to this Run.",
        });
      }
    },
    1000,
    10000
  );

  const startRun = async () => {
    console.log("start run");
    try {
      const input = {
        ID: run.id,
        projectID: project.id,
      };

      console.log(JSON.stringify(input, null, "  "));

      await mutate(client, {
        mutation: START_RUN,
        variables: {
          input,
        },
      });

      notify({
        success: true,
        title: `Run Started`,
        // body:
        //   "Something happened on the server, and we could not create a new Run as requested.",
      });
    } catch (error) {
      console.error(error);
      notify({
        failed: true,
        title: `Could not start Run`,
        body:
          "Something happened on the server, and we could not start the Run.",
      });
    }
  };

  let template = deepCopy(run.template);

  function handleClick(event) {
    const { action } = event.detail;
    switch (action) {
      case "start":
        startRun();
        break;
      default:
        break;
    }
  }

  let actions = [];
  let facts = [];
  let disabled = false;
  $: isMturkQual = template.selectionType === "MTURK_QUALIFICATIONS";

  $: {
    if (
      template.steps.length === 0 ||
      (isMturkQual && template.steps[0].type !== "MTURK_HIT")
    ) {
      disabled = true;
    } else {
      disabled = false;
    }
  }

  $: {
    actions = [];
    facts = [];

    actions.push({
      text: "Duplicate",
      action: "duplicate",
      icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M464 0H144c-26.51 0-48 21.49-48 48v48H48c-26.51 0-48 21.49-48 48v320c0 26.51 21.49 48 48 48h320c26.51 0 48-21.49 48-48v-48h48c26.51 0 48-21.49 48-48V48c0-26.51-21.49-48-48-48zM362 464H54a6 6 0 0 1-6-6V150a6 6 0 0 1 6-6h42v224c0 26.51 21.49 48 48 48h224v42a6 6 0 0 1-6 6zm96-96H150a6 6 0 0 1-6-6V54a6 6 0 0 1 6-6h308a6 6 0 0 1 6 6v308a6 6 0 0 1-6 6z"/></svg>`,
      primary: false,
    });

    switch (run.status) {
      case "CREATED":
      case "DONE":
      case "TERMINATED":
      case "FAILED":
        actions.push({
          text: "Archive",
          action: "archive",
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M464 32H48C21.5 32 0 53.5 0 80v80c0 8.8 7.2 16 16 16h16v272c0 17.7 14.3 32 32 32h384c17.7 0 32-14.3 32-32V176h16c8.8 0 16-7.2 16-16V80c0-26.5-21.5-48-48-48zm-32 400H80V176h352v256zm32-304H48V80h416v48zM204 272h104c6.6 0 12-5.4 12-12v-24c0-6.6-5.4-12-12-12H204c-6.6 0-12 5.4-12 12v24c0 6.6 5.4 12 12 12z"/></svg>`,
          primary: false,
        });
        break;
      default:
        break;
    }

    switch (run.status) {
      case "CREATED":
        actions.push({
          text: "Schedule",
          action: "schedule",
          disabled,
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M256 8C119 8 8 119 8 256s111 248 248 248 248-111 248-248S393 8 256 8zm0 448c-110.5 0-200-89.5-200-200S145.5 56 256 56s200 89.5 200 200-89.5 200-200 200zm61.8-104.4l-84.9-61.7c-3.1-2.3-4.9-5.9-4.9-9.7V116c0-6.6 5.4-12 12-12h32c6.6 0 12 5.4 12 12v141.7l66.8 48.6c5.4 3.9 6.5 11.4 2.6 16.8L334.6 349c-3.9 5.3-11.4 6.5-16.8 2.6z"/></svg>`,
          primary: false,
        });
        actions.push({
          text: "Start Now",
          action: "start",
          disabled,
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M424.4 214.7L72.4 6.6C43.8-10.3 0 6.1 0 47.9V464c0 37.5 40.7 60.1 72.4 41.3l352-208c31.4-18.5 31.5-64.1 0-82.6zM48 453.5v-395c0-4.6 5.1-7.5 9.1-5.2l334.2 197.5c3.9 2.3 3.9 8 0 10.3L57.1 458.7c-4 2.3-9.1-.6-9.1-5.2z"/></svg>`,
          primary: true,
        });

        if (run.startAt) {
          facts.push({
            text: `Starts ${dayjs(run.startAt).calendar()}`,
            icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M256 8C119 8 8 119 8 256s111 248 248 248 248-111 248-248S393 8 256 8zm0 448c-110.5 0-200-89.5-200-200S145.5 56 256 56s200 89.5 200 200-89.5 200-200 200zm61.8-104.4l-84.9-61.7c-3.1-2.3-4.9-5.9-4.9-9.7V116c0-6.6 5.4-12 12-12h32c6.6 0 12 5.4 12 12v141.7l66.8 48.6c5.4 3.9 6.5 11.4 2.6 16.8L334.6 349c-3.9 5.3-11.4 6.5-16.8 2.6z"/></svg>`,
          });
        } else {
          facts.push({
            text: `Not scheduled`,
            icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M256 8C119.043 8 8 119.083 8 256c0 136.997 111.043 248 248 248s248-111.003 248-248C504 119.083 392.957 8 256 8zm0 448c-110.532 0-200-89.431-200-200 0-110.495 89.472-200 200-200 110.491 0 200 89.471 200 200 0 110.53-89.431 200-200 200zm107.244-255.2c0 67.052-72.421 68.084-72.421 92.863V300c0 6.627-5.373 12-12 12h-45.647c-6.627 0-12-5.373-12-12v-8.659c0-35.745 27.1-50.034 47.579-61.516 17.561-9.845 28.324-16.541 28.324-29.579 0-17.246-21.999-28.693-39.784-28.693-23.189 0-33.894 10.977-48.942 29.969-4.057 5.12-11.46 6.071-16.666 2.124l-27.824-21.098c-5.107-3.872-6.251-11.066-2.644-16.363C184.846 131.491 214.94 112 261.794 112c49.071 0 101.45 38.304 101.45 88.8zM298 368c0 23.159-18.841 42-42 42s-42-18.841-42-42 18.841-42 42-42 42 18.841 42 42z"/></svg>`,
          });
        }
        break;
      case "running":
        actions.push({
          text: "Cancel",
          action: "cancel",
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512"><path d="M207.6 256l107.72-107.72c6.23-6.23 6.23-16.34 0-22.58l-25.03-25.03c-6.23-6.23-16.34-6.23-22.58 0L160 208.4 52.28 100.68c-6.23-6.23-16.34-6.23-22.58 0L4.68 125.7c-6.23 6.23-6.23 16.34 0 22.58L112.4 256 4.68 363.72c-6.23 6.23-6.23 16.34 0 22.58l25.03 25.03c6.23 6.23 16.34 6.23 22.58 0L160 303.6l107.72 107.72c6.23 6.23 16.34 6.23 22.58 0l25.03-25.03c6.23-6.23 6.23-16.34 0-22.58L207.6 256z"/></svg>`,
          primary: true,
        });

        facts.push({
          text: `Started at ${dayjs(run.startedAt).calendar()}`,
          icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M256 8C119 8 8 119 8 256s111 248 248 248 248-111 248-248S393 8 256 8zm0 448c-110.5 0-200-89.5-200-200S145.5 56 256 56s200 89.5 200 200-89.5 200-200 200zm61.8-104.4l-84.9-61.7c-3.1-2.3-4.9-5.9-4.9-9.7V116c0-6.6 5.4-12 12-12h32c6.6 0 12 5.4 12 12v141.7l66.8 48.6c5.4 3.9 6.5 11.4 2.6 16.8L334.6 349c-3.9 5.3-11.4 6.5-16.8 2.6z"/></svg>`,
        });
        break;
      case "DONE":
      case "TERMINATED":
      case "FAILED":
        let verb = "";

        switch (run.status) {
          case "done":
            verb = "Finished";
            break;
          case "terminated":
            verb = "Cancelld";
            break;
          case "failed":
            verb = "Failed";
            break;
          default:
            console.error("unknown status", run.status);
            break;
        }

        if (run.endedAt) {
          facts.push({
            text: `${verb} at ${dayjs(run.endedAt).calendar()}`,
            icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M256 8C119 8 8 119 8 256s111 248 248 248 248-111 248-248S393 8 256 8zm0 448c-110.5 0-200-89.5-200-200S145.5 56 256 56s200 89.5 200 200-89.5 200-200 200zm61.8-104.4l-84.9-61.7c-3.1-2.3-4.9-5.9-4.9-9.7V116c0-6.6 5.4-12 12-12h32c6.6 0 12 5.4 12 12v141.7l66.8 48.6c5.4 3.9 6.5 11.4 2.6 16.8L334.6 349c-3.9 5.3-11.4 6.5-16.8 2.6z"/></svg>`,
          });

          if (run.startedAt) {
            const from = dayjs(run.startedAt);
            const to = dayjs(run.endedAt);
            const relative = from.to(to, true);
            const duration = dayjs.duration(to.diff(from));
            const minutes = Math.round(duration.as("minute"));
            const human = minutes + ` minute${minutes == 1 ? "" : "s"}`;
            const output = minutes < 120 ? human : relative;

            facts.push({
              text: `Lasted at ${output}`,
              icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M393.9 184l22.6-22.6c4.7-4.7 4.7-12.3 0-17l-17-17c-4.7-4.7-12.3-4.7-17 0l-20.7 20.7c-31.1-27.5-70.4-45.9-113.8-50.8V48h28c6.6 0 12-5.4 12-12V12c0-6.6-5.4-12-12-12H172c-6.6 0-12 5.4-12 12v24c0 6.6 5.4 12 12 12h28v49.4C96.4 109.3 16 197.2 16 304c0 114.9 93.1 208 208 208s208-93.1 208-208c0-44.7-14.1-86.1-38.1-120zM224 464c-88.4 0-160-71.6-160-160s71.6-160 160-160 160 71.6 160 160-71.6 160-160 160zm12-112h-24c-6.6 0-12-5.4-12-12V204c0-6.6 5.4-12 12-12h24c6.6 0 12 5.4 12 12v136c0 6.6-5.4 12-12 12z"/></svg>`,
            });
          }
        }

      default:
        break;
    }
  }
</script>

{#if run}
  <Layout
    bind:title={name}
    on:click={handleClick}
    {actions}
    {facts}
    titleUpdatable>
    <div slot="pretitle" class="h-full flex items-center">
      <StatusBadge large status={run.status} />
    </div>

    {#if run.status === 'RUNNING'}
      <RunningRun {project} {run} />
    {:else}
      <Template {project} {run} bind:template />
    {/if}
  </Layout>
{/if}
