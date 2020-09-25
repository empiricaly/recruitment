import App from "./App.svelte";
import { manageFocus } from "./components/base/focus.js";
import "./inits/dayjs.js";
import "./inits/routes.js";
import "./inits/shortcuts.js";

manageFocus();

const app = new App({
  target: document.body,
});

export default app;

// recreate the whole app if an HMR update touches this module
if (import.meta.hot) {
  import.meta.hot.dispose(() => {
    app.$destroy();
  });
  import.meta.hot.accept();
}
