const columnDelimiter = ",";
const lineDelimiter = "\r\n";
const doubleQuote = '"';
const doubleDoubleQuote = '""';

// Loosely https://www.ietf.org/rfc/rfc4180.txt
export function toCSV(data) {
  if (data === null || data.length === 0) {
    return "";
  }

  // Expecting all objects have all the keys (bad idea...)
  const keys = Object.keys(data[0]);

  // Header
  let result = keys.join(columnDelimiter) + lineDelimiter;

  for (let i = 0; i < data.length; i++) {
    const values = keys.map((key) => {
      let value = data[i][key];

      if (typeof value !== "string") {
        return JSON.stringify(value);
      }

      const quoted = value.includes(doubleQuote);
      if (
        value.includes(columnDelimiter) ||
        value.includes(lineDelimiter) ||
        quoted
      ) {
        if (quoted) {
          value = value.replaceAll(doubleQuote, doubleDoubleQuote);
        }
        return `${doubleQuote}${value}${doubleQuote}`;
      } else {
        return value;
      }
    });

    result += values.join(columnDelimiter) + lineDelimiter;
  }

  return result;
}
