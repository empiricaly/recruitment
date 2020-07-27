export function uniqueID() {
  const uniqID = genID();
  return (k) => `${uniqID}_${k}`;
}

export function genID() {
  return Math.random().toString(36).substring(7);
}
