const debug = false;
import { pathToRegexp } from "path-to-regexp";
import { init } from "../lib/routing.js";
import AllParticipants from "../pages/AllParticipants.svelte";
import Overview from "../pages/Overview.svelte";
import Participants from "../pages/Participants.svelte";
import Profile from "../pages/Profile.svelte";
import Projects from "../pages/Projects.svelte";
import Run from "../pages/Run.svelte";
import Runs from "../pages/Runs.svelte";
import Signin from "../pages/Signin.svelte";
import Template from "../pages/Template.svelte";
import Templates from "../pages/Templates.svelte";

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

init({
  mode: "history",
  routes: [
    {
      path: signinPath,
      component: Signin,
    },
    {
      path: "/projects",
      component: Projects,
    },
    {
      path: "/projects/:projectID/overview",
      component: Overview,
    },
    {
      path: "/projects/:projectID/runs",
      component: Runs,
    },
    {
      path: "/projects/:projectID/runs/:runID",
      component: Run,
    },
    {
      path: "/projects/:projectID/templates",
      component: Templates,
    },
    {
      path: "/projects/:projectID/templates/:templateID",
      component: Template,
    },
    {
      path: "/projects/:projectID/participants",
      component: Participants,
    },
    {
      path: "/participants",
      component: AllParticipants,
    },
    {
      path: "/profile",
      component: Profile,
    },
  ],
});
