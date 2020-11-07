export function participantPerQueryType(type, result) {
  switch (type) {
    case "run":
      const steps = result.data.project.runs[0].steps;
      if (steps.length === 0) {
        return;
      }
      return steps[0].participants;
    case "project":
      return result.data.project.participants;
    case "all":
      return result.data.participants;
    default:
      console.error("unknown type");
  }
}

export function participantsExportFormat(participants, keys, flat = false) {
  const out = [];
  for (let i = 0; i < participants.length; i++) {
    const { id, mturkWorkerID, createdAt, data } = participants[i];
    const p = { id, mturkWorkerID, createdAt };
    if (!flat) {
      p.data = {};
    }

    for (let j = 0; j < keys.length; j++) {
      const d = data.find((d) => d.key === keys[j]);
      if (flat) {
        p[`data_${keys[j]}`] = d ? d.val : null;
      } else {
        if (d) {
          p.data[keys[j]] = d.val;
        }
      }
    }
    out.push(p);
  }
  return out;
}
