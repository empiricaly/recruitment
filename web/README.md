# Empirica Recruitment Web front

The web front for Empirica Recruitment is a SPA (Single-Page Application),
meaning it is statically generated and served. The app then connects to the
GraphQL backend to fetch data.

## Build

To build the static export of the application, simply run:

```sh
yarn build
```

It will export everything needed to run the application in `/public`.

## Solo start

By default, the web development server will be started automatically at the same
time as the GraphQL server backend, thanks to modd (see root Readme). If you
just want to start the web app in dev mode manually, run:

```sh
yarn dev
```

Navigate to [localhost:5000](http://localhost:5000). You should see your app
running.

**:warning: Experimental :warning:**

`yarn dev` (or `npm run dev`) will start an HMR Nollup dev build, which is
dramatically faster than rollup + livereload, but also still experimental (+ HMR
can sometimes not reinitialize a component correctly). So just reload if you
have any issues, or switch over to `yarn dev:livereload` if you are having
issues.
