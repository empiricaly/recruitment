import { InMemoryCache } from "@apollo/client/cache";
import { ApolloClient, split } from "@apollo/client/core";
import { setContext } from "@apollo/client/link/context";
import { HttpLink } from "@apollo/client/link/http";
import { WebSocketLink } from "@apollo/client/link/ws";
import { getMainDefinition } from "@apollo/client/utilities";

export const uri =
  process.env.NODE_ENV === "development"
    ? "http://localhost:8880/query"
    : document.location.origin + "/query";

// Inject authentication
const authLink = setContext((_, { headers }) => {
  // get the authentication token from local storage if it exists
  const token = localStorage.getItem("token");
  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    },
    fetchOptions: {
      mode: "cors",
    },
  };
});

// Create an HTTP link:
const httpLink = new HttpLink({
  uri, // use https for secure endpoint
  // credentials: "include",
  fetchOptions: {
    mode: "cors",
  },
});

// Create a WebSocket link
const wsLink = new WebSocketLink({
  uri: uri.replace("http", "ws"), // use wss for a secure endpoint
  // credentials: "include",
  options: {
    reconnect: true,
    fetchOptions: {
      mode: "cors",
    },
  },
});

// Using the ability to split links, you can send data to each link
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

export const client = new ApolloClient({
  link: authLink.concat(link),
  cache: new InMemoryCache(),
});
