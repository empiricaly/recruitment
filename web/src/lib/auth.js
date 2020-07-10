import { writable } from "svelte/store";

const { subscribe: subscribeLoggedIn, set: setLoggedIn } = writable(false);
export const loggingIn = { subscribe: subscribeLoggedIn, set: setLoggedIn };

const { subscribe: subscribeUser, set: setUser } = writable(false);
export const user = { subscribe: subscribeUser, set: setUser };
