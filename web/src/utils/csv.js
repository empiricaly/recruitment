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
        if (value === null || value === undefined) {
          return "";
        }
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

export function fromCSVToJSON(csv) {
  const lines = csv.split(lineDelimiter);
  const result = [];
  const headers = lines[0].split(columnDelimiter);

  for (let i = 1; i < lines.length; i++) {
    if (!lines[i]) continue;
    const obj = {};
    const currentline = lines[i].split(columnDelimiter);

    for (let j = 0; j < headers.length; j++) {
      let cell = currentline[j];

      if (cell.includes(doubleDoubleQuote)) {
        cell = cell.replaceAll(doubleDoubleQuote, "");
      }

      obj[headers[j]] = cell;
    }
    result.push(obj);
  }
  return result;
}
