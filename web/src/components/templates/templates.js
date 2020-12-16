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
  subject: "",
  url:
    "function (participant, run, currentStep, steps) {\n\treturn `https://example.com`;\n}",
  message: "\n\n\n\n",
  messageType: "PLAIN",
  lobby: "",
  lobbyType: "PLAIN",
  lobbyExpiration: 0,
};

export const defaultHITMessageStepArgs = {
  subject: "",
  url:
    "function (participant, run, currentStep, steps) {\n\treturn `https://example.com`;\n}",
  message: "\n\n\n\n",
  messageType: "HTML",
  lobby: "",
  lobbyType: "HTML",
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
  js: "function (participants) {\n\treturn participants;\n}",
  filter: "",
};
