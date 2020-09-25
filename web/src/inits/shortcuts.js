import hotkeys from "hotkeys-js";
import { logout } from "../lib/auth.js";

// Logout
hotkeys("ctrl+e", function (event, handler) {
  logout();
});
