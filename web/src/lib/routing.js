import { createBrowserHistory, createHashHistory } from "history";
import { match, pathToRegexp } from "path-to-regexp";
import { writable } from "svelte/store";

export const path = writable("/");
export const params = writable({});

class Router {
  constructor({ target, mode = "hash", routes = [] }) {
    let history;
    switch (mode) {
      case "hash":
        history = createHashHistory();
        break;
      case "history":
        history = createBrowserHistory();
        break;
      default:
        history = createHashHistory();
        break;
    }

    path.set(history.location.pathname);
    history.listen(({ pathname }) => {
      path.set(pathname);
    });

    Router.mode = mode;
    Router.history = history;
    Router.push = (path) => history.push(path);
    Router.replace = (path) => history.replace(path);
    Router.go = (n) => history.go(n);
    Router.goBack = () => history.goBack();
    Router.goForward = () => history.goForward();
    Router.listen = (fn) => history.listen(fn);

    this.$content = null;
    this.target =
      typeof target === "string" ? document.querySelector(target) : target;
    this.routes = routes;
    this.$listener = history.listen(this.handleRouteChange.bind(this));

    this.handleRouteChange(history.location);
  }

  destroy() {
    if (this.$listener) {
      this.$listener();
      this.$listener = null;
    }
  }

  handleRouteChange({ pathname }) {
    let matchedRoute;
    let prms;

    for (const route of this.routes) {
      const regexp = pathToRegexp(route.path);
      if (regexp.test(pathname)) {
        const m = match(route.path, { decode: decodeURIComponent });
        prms = m(pathname).params;
        params.set(prms);
        matchedRoute = route;
        break;
      }
    }

    if (matchedRoute && matchedRoute.component) {
      if (this.$content) this.$content.$destroy();
      const { component: Component, props } = matchedRoute;

      this.$content = new Component({
        target: this.target,
        props: { params: prms, ...props },
      });
    } else {
      params.set({});
    }
  }
}

export default Router;
