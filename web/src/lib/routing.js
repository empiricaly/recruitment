import { createBrowserHistory } from "history";
import { match, pathToRegexp } from "path-to-regexp";
import { writable } from "svelte/store";

export const path = writable("/");
export const params = writable({});

const history = createBrowserHistory();
path.set(history.location.pathname);
history.listen(({ pathname }) => {
  path.set(pathname);
});

let rroutes;
export function init({ routes }) {
  rroutes = routes;
  if (handleRouteChange) {
    handleRouteChange(history.location);
  }
}

export const getPath = () => history.location.pathname;
export const push = (path) => history.push(path);
export const replace = (path) => history.replace(path);
export const go = (n) => history.go(n);
export const goBack = () => history.goBack();
export const goForward = () => history.goForward();
export const listen = (fn) => history.listen(fn);

let handleRouteChange;
export function route(target) {
  let content = null;
  handleRouteChange = function ({ pathname }) {
    let matchedRoute;

    for (const route of rroutes) {
      const regexp = pathToRegexp(route.path);
      if (regexp.test(pathname)) {
        const m = match(route.path, { decode: decodeURIComponent });
        const prms = m(pathname).params;
        params.set(prms);
        matchedRoute = route;
        break;
      }
    }

    if (matchedRoute && matchedRoute.component) {
      if (content) content.$destroy();
      const { component: Component, props } = matchedRoute;

      content = new Component({
        target,
        props: { ...props },
      });
    } else {
      params.set({});
    }
  };

  let listener = history.listen(handleRouteChange);
  handleRouteChange(history.location);

  function destroy() {
    if (!listener) {
      return;
    }

    listener();
    listener = null;
  }

  return { destroy };
}
