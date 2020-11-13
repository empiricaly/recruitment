export function isFloat(value) {
  let splittedValues = value.split(".");

  if (splittedValues.length !== 2) {
    return false;
  }

  for (let i = 0; i < splittedValues.length; i++) {
    if (!Number.isInteger(+splittedValues[i])) {
      return false;
    }
  }

  return true;
}

export function isInteger(value) {
  return Number.isInteger(+value);
}

export function isBoolean(value) {
  return value === "true" || value === "false";
}
