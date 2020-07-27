import { InMemoryCache } from "apollo-cache-inmemory";
import ApolloClient from "apollo-client";
import { split } from "apollo-link";
import { createHttpLink } from "apollo-link-http";
import { WebSocketLink } from "apollo-link-ws";
import { getMainDefinition } from "apollo-utilities";
import { writable } from "svelte/store";
export const uri =
  process.env.NODE_ENV === "development"
    ? "http://localhost:8880/query"
    : document.location.origin + "/query";
export const client = writable(null);
const httpLink = createHttpLink({
  uri, // use https for secure endpoint
  fetchOptions: {
    mode: "cors",
  },
});
// Create a WebSocket link:
const wsLink = new WebSocketLink({
  uri: uri.replace("http", "ws"), // use wss for a secure endpoint
  options: {
    reconnect: true,
    fetchOptions: {
      mode: "cors",
    },
  },
});
// using the ability to split links, you can send data to each link
// depending on what kind of operation is being sent
const link = split(
  // split based on operation type
  ({ query }) => {
    const { kind, operation } = getMainDefinition(query);
    return kind === "OperationDefinition" && operation === "subscription";
  },
  wsLink,
  httpLink
);
client.set(
  new ApolloClient({
    link,
    cache: new InMemoryCache(),
  })
);
