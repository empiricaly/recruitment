import { mutate } from "svelte-apollo";
import { get, writable } from "svelte/store";
import { defaultPath, isPublicPath, signinPath } from "../inits/routeparts.js";
import { client } from "../lib/apollo.js";
import { AUTH, ME } from "../lib/queries.js";
import { getPath, listen, replace } from "./routing.js";

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
    try {
      // query user
      // if fail to get user => token invalid?
      const meResp = await client.query({ query: ME });
      const me = meResp.data.me;
      console.log("me", me);
      setUser(me.__typename !== "Admin" ? null : me);
      authRedirect();
    } catch (error) {
      console.log(error);
      localStorage.removeItem("token");
      setToken(null);
    }
  } else {
    setUser(null);
    authRedirect();
  }
  setLoggingIn(false);
});

function authRedirect() {
  const path = getPath();
  const loggedIn = get(user);
  if (loggedIn && path === signinPath) {
    replace(initialPath !== path ? initialPath : defaultPath);
  } else if (!loggedIn && !isPublicPath(path)) {
    replace(signinPath);
  }
}
listen(authRedirect);
