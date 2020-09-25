import { mutate } from "svelte-apollo";
import { writable } from "svelte/store";
import { client } from "../lib/apollo";
import { AUTH, ME } from "../lib/queries";
import { getPath, push, replace } from "./routing.js";

const { subscribe: subscribeLoggingIn, set: setLoggingIn } = writable(true);
export const loggingIn = { subscribe: subscribeLoggingIn, set: setLoggingIn };

const { subscribe: subscribeUser, set: setUser } = writable(false);
export const user = { subscribe: subscribeUser, set: setUser };

const { subscribe: subscribeToken, set: setToken } = writable(
  localStorage.getItem("token")
);
export const token = { subscribe: subscribeToken, set: setToken };

export const authenticate = async (username, password) => {
  const authResp = await mutate(client, {
    mutation: AUTH,
    variables: { input: { username, password } },
  });
  localStorage.setItem("token", authResp.data.auth.token);
  setToken(authResp.data.auth.token);
};

export function logout() {
  localStorage.removeItem("token");
  setToken(null);
}

let initialPath = document.location.pathname;
subscribeToken(async (token) => {
  console.log("token:", token);
  if (token) {
    // query user
    // if fail to get user => token invalid?
    const meResp = await client.query({ query: ME });
    const me = meResp.data.me;
    console.log("me", me);
    setUser(me.__typename !== "Admin" ? null : me);
    push(initialPath !== getPath() ? initialPath : "/projects");
  } else {
    setUser(null);
    replace("/signin");
  }
  setLoggingIn(false);
});

// function authRedirect() {
//   const path = getPath();
//   const u = get(user);
//   if (u && path === signinPath) {
//     replace(initialPath !== path ? initialPath : "/projects");
//   } else if (!u && !isPublicPath(path)) {
//     replace(signinPath);
//   }
// }
// listen(authRedirect);
// authRedirect();
