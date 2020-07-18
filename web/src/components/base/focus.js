const TAB_KEY_CODE = 9;

const el = document.documentElement;

export function manageFocus() {
  el.addEventListener("mousedown", handleMouseDown);
}

function reset() {
  el.classList.remove("focus-disabled");
  el.removeEventListener("keydown", handleKeyDown);
  el.removeEventListener("mousedown", handleMouseDown);
}

function handleMouseDown(event) {
  reset();
  el.classList.add("focus-disabled");
  el.addEventListener("keydown", handleKeyDown);
}

function handleKeyDown(event) {
  if (event.which === TAB_KEY_CODE) {
    reset();
    el.addEventListener("mousedown", handleMouseDown);
  }
}
