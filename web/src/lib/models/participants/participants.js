import dayjs from "dayjs";
import { fromCSVToJSON } from "../../../utils/csv";
import { isBoolean, isFloat, isInteger } from "../../../utils/typeValue.js";
import { download } from "../../../utils/download.js";
import { toCSV } from "../../../utils/csv.js";

export function setValue(value) {
  value = value.trim();
  if (value === "") {
    return;
  }

  if (isFloat(value)) {
    return parseFloat(value);
  } else if (isInteger(value)) {
    return parseInt(value);
  } else if (isBoolean(value)) {
    return value === "true";
  } else {
    return value;
  }
}

export function participantPerQueryType(type, result) {
  switch (type) {
    case "run":
      const steps = result.data.project.runs[0].steps;
      if (steps.length === 0) {
        return;
      }
      return {
        participants: steps[0].participants,
        total: steps[0].participantsCount,
      };
    case "project":
      return {
        participants: result.data.project.participants,
        total: result.data.project.participantsCount,
      };
    case "all":
      return {
        participants: result.data.participants,
        total: result.data.participantCount,
      };
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

export function getParticipants(file, customData, callback) {
  const r = new FileReader();
  let participants = [];
  r.onload = (e) => {
    const text = e.target.result;
    switch (file.type) {
      case "text/csv":
        participants = fromCSVToJSON(text);
        participants = participants.map((p) => {
          let data = {};
          for (const key in p) {
            const splittedKey = key.split("_");
            if (splittedKey[0] === "data") {
              splittedKey.shift();
              const newKey = splittedKey.join("_");
              data = { ...data, [newKey.trim()]: setValue(p[key]) };
            }
          }
          return {
            id: p.id,
            mturkWorkerID: p.mturkWorkerID,
            createdAt: p.createdAt,
            data,
          };
        });
        break;
      case "application/json":
        participants = JSON.parse(text);
        break;

      default:
        return callback(null, "Unsupported type of file.");
    }

    participants = participants
      .map(({ id, mturkWorkerID, data, createdAt }) => {
        data = { ...data, ...customData };
        // return { id, mturkWorkerID, data: JSON.stringify(data), createdAt };
        return { id, mturkWorkerID, data, createdAt };
      })
      .map((p) => {
        const data = [];
        for (const key in p.data) {
          data.push({ key, val: p.data[key] + "" });
        }
        return { ...p, data };
      });
    return callback(participants, null);
  };
  r.readAsText(file);
}

export function exportJson(participants, keys) {
  const out = participantsExportFormat(participants, keys);
  const content = JSON.stringify(out);
  const mime = "application/json;charset=utf-8";
  const date = dayjs().format("YYYY-MM-DDTHH:mm:ss");
  const filename = `Empirica recruitment export – ${date}.json`;
  download(content, filename, mime);
}

export function exportCSV(participants, keys) {
  const out = participantsExportFormat(participants, keys, true);
  const content = toCSV(out);
  const mime = "text/csv;charset=utf-8";
  const date = dayjs().format("YYYY-MM-DDTHH:mm:ss");
  const filename = `Empirica recruitment export – ${date}.csv`;
  download(content, filename, mime);
}
