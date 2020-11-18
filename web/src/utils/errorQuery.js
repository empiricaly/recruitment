export function handleErrorMessage({ graphQLErrors, networkError }) {
  if (graphQLErrors) {
    graphQLErrors.forEach(({ message, locations, path }) =>
      console.error(
        `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
      )
    );
  }

  if (networkError) {
    const {
      name,
      message,
      statusCode,
      result: { errors },
    } = networkError;
    console.error(
      `[Network error]: ${name}, status code: ${statusCode}, Message: ${message} `
    );
    errors.forEach(({ message, path, extensions }) => {
      console.error(
        `[Network error]: errors: Message: ${message}, path: "${path}", extensions: ${extensions.code} `
      );
    });
  }
}
