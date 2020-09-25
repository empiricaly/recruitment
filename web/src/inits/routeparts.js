import { pathToRegexp } from "path-to-regexp";

const debug = false;

export const signinPath = "/signin";
export const defaultPath = "/projects";
export const publicPaths = [signinPath, "/", "/lobby/:id"];
const publicPathsRegexp = [];

for (const path of publicPaths) {
  publicPathsRegexp.push(pathToRegexp(path));
}

export function isPublicPath(path) {
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
