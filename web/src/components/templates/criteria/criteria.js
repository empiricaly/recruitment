export const comparators = [
  {
    label: "equal to",
    value: "EQUAL_TO",
  },
  {
    label: "not equal to",
    value: "NOT_EQUAL_TO",
  },
  {
    label: "<",
    value: "LESS_THAN",
    title: "lesser than",
  },
  {
    label: "≤",
    value: "LESS_THAN_OR_EQUAL_TO",
    title: "lesser than or equal to",
  },
  {
    label: ">",
    value: "GREATER_THAN",
    title: "greater than",
  },
  {
    label: "≥",
    value: "GREATER_THAN_OR_EQUAL_TO",
    title: "greater than or equal to",
  },
  {
    label: "exists",
    value: "EXISTS",
  },
  {
    label: "does not exist",
    value: "DOES_NOT_EXIST",
  },
  {
    label: "in",
    value: "IN",
  },
  {
    label: "not in",
    value: "NOT_IN",
  },
];

const comparatorsIndex = {};
for (const comparator of comparators) {
  comparatorsIndex[comparator.value] = comparator.title || comparator.label;
}
export { comparatorsIndex };
