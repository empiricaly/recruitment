export const selectionTypes = [
  {
    label: "Internal Database",
    value: "INTERNAL_DB",
  },
  {
    label: "MTurk Qualifications",
    value: "MTURK_QUALIFICATIONS",
  },
];

export const defaultMessageStepArgs = {
  url: "",
  message: "\n\n\n\n",
  messageType: "MARKDOWN",
  lobby: "",
  lobbyType: "MARKDOWN",
  lobbyExpiration: 0,
};

export const defaultHITStepArgs = {
  title: "",
  description: "",
  keywords: "",
  microbatch: false,
  reward: null,
  timeout: 0,
  duration: 60,
  workersCount: 0,
};

export const defaultFilterStepArgs = {
  type: "JS",
  js: "function (players, step, stepRun) {\n\treturn players;\n}",
  filter: "",
};
