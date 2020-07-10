<script context="module">
  import { pathToRegexp } from "path-to-regexp";

  const signinPath = "/signin";
  const publicPaths = [signinPath, "/lobby/:id"];
  const publicPathsRegexp = [];

  for (const path of publicPaths) {
    publicPathsRegexp.push(pathToRegexp(path));
  }

  function isPublicPath(path) {
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
  import Home from "./pages/Home.svelte";
  import Signin from "./pages/Signin.svelte";
  import Link from "./components/base/Link.svelte";

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
          path: "/",
          component: Home
        },
        {
          path: signinPath,
          component: Signin
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

<style>
  main {
    width: 100vw;
    height: 100vh;
  }
</style>
