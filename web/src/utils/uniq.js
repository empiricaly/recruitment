export function uniqueID() {
  const uniqID = Math.random().toString(36).substring(7);
  return (k) => `${uniqID}_${k}`;
}
