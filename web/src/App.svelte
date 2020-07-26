<script context="module">
  const debug = true;
  import { pathToRegexp } from "path-to-regexp";

  const signinPath = "/signin";
  const publicPaths = [signinPath, "/", "/lobby/:id"];
  const publicPathsRegexp = [];

  for (const path of publicPaths) {
    publicPathsRegexp.push(pathToRegexp(path));
  }

  function isPublicPath(path) {
    if (debug) {
      return true;
    }
    for (const regexp of publicPathsRegexp) {
      if (regexp.test(path)) {
        return true;
      }
    }
    return false; // yo
  }
</script>

<script>
  import Router from "./lib/routing.js";
  import { user } from "./lib/auth.js";
  import Projects from "./pages/Projects.svelte";
  import Overview from "./pages/Overview.svelte";
  import Runs from "./pages/Runs.svelte";
  import Run from "./pages/Run.svelte";
  import Template from "./pages/Template.svelte";
  import Templates from "./pages/Templates.svelte";
  import Participants from "./pages/Participants.svelte";
  import AllParticipants from "./pages/AllParticipants.svelte";
  import Profile from "./pages/Profile.svelte";
  import Signin from "./pages/Signin.svelte";
  import Alert from "./components/overlays/Alert.svelte";

  let initialPath = document.location.pathname;
  function authRedirect() {
    const path = Router.history.location.pathname;
    if ($user && path === signinPath) {
      Router.replace(initialPath !== path ? initialPath : "/");
    } else if (!$user && !isPublicPath(path)) {
      Router.replace(signinPath);
    }
  }

  function create(node) {
    const router = new Router({
      target: node,
      mode: "history",
      routes: [
        {
          path: signinPath,
          component: Signin
        },
        {
          path: "/projects",
          component: Projects
        },
        {
          path: "/projects/:projectID/overview",
          component: Overview
        },
        {
          path: "/projects/:projectID/runs",
          component: Runs
        },
        {
          path: "/projects/:projectID/runs/:runID",
          component: Run
        },
        {
          path: "/projects/:projectID/templates",
          component: Templates
        },
        {
          path: "/projects/:projectID/templates/:templateID",
          component: Template
        },
        {
          path: "/projects/:projectID/participants",
          component: Participants
        },
        {
          path: "/participants",
          component: AllParticipants
        },
        {
          path: "/profile",
          component: Profile
        }
      ]
    });

    Router.listen(authRedirect);
    authRedirect();

    return {
      destroy() {
        router.destroy();
      }
    };
  }
</script>

<main class="antialiased bg-gray-100" use:create />
<Alert />

<style>
  main {
    width: 100vw;
    height: 100vh;
  }
</style>
