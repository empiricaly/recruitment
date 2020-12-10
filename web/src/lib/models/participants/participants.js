import dayjs from "dayjs";
import { query, mutate } from "svelte-apollo";

import { client } from "../../apollo";
import {
  GET_PROJECT_PARTICIPANTS,
  GET_ALL_PARTICIPANTS,
  ADD_PARTICIPANTS,
} from "../../queries";
import { fromCSVToJSON } from "../../../utils/csv";
import { isBoolean, isFloat, isInteger } from "../../../utils/typeValue.js";
import { download } from "../../../utils/download.js";
import { toCSV } from "../../../utils/csv.js";
import { handleErrorMessage } from "../../../utils/errorQuery.js";
import { notify } from "../../../components/overlays/Notification.svelte";

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

export async function fetchParticipants(
  all = false,
  project,
  type = "json",
  keys = {},
  setLoading
) {
  const limit = 100;
  let offset = 0;
  let args = {
    query: !all ? GET_PROJECT_PARTICIPANTS : GET_ALL_PARTICIPANTS,
    variables: {
      offset,
      limit,
    },
  };

  if (project) {
    args.variables.projectID = project.projectID;
  }

  try {
    let finish = false;
    let allParticipants = [];

    // fetching all participants by paginating them
    while (!finish) {
      const participantsQuery = query(client, args);
      const result = await participantsQuery.refetch();
      const pp = participantPerQueryType(!all ? "project" : "all", result);
      if (pp) {
        if (!pp.participants || pp.participants.length === 0) {
          finish = true;
          continue;
        }

        allParticipants = allParticipants.concat(pp.participants);
        offset++;
        args.variables.offset = offset;
      } else {
        finish = true;
      }
    }

    switch (type) {
      case "json":
        exportJson(allParticipants, keys);
        break;

      case "csv":
        exportCSV(allParticipants, keys);
        break;

      default:
        console.error("unknown file type");
        notify({
          failed: true,
          title: `Could not export participants. Unknown file type.`,
        });
        setLoading(false);
        return;
    }
    notify({
      failed: false,
      title: `Participants exported successfully.`,
    });
    setLoading(false);
  } catch (error) {
    handleErrorMessage(error);
    notify({
      failed: true,
      title: `Could not export participants.`,
    });
    setLoading(false);
  }
}

export async function exportParticipants({
  all = false,
  project,
  type = "json",
  keys,
  setLoading,
}) {
  let worker = new Worker(
    `data:text/javascript,
    onmessage = async function(event){    
      ${await fetchParticipants(all, project, type, keys, setLoading)}
    };
    `
  );

  worker.postMessage({});
}

export function importParticipants({
  files,
  customKey,
  customValue,
  dispatch,
  projectID,
  setOpen,
}) {
  let file = files.length > 0 ? files[0] : null;
  let isCustomEmpty;
  let customData = {};

  if (!file) {
    notify({
      failed: true,
      title: `Could not import participants.`,
      body: "No file selected.",
    });
    return;
  }

  if (customKey || customValue) {
    if (
      !customKey ||
      !customKey.trim() ||
      !customValue ||
      !customValue.trim()
    ) {
      isCustomEmpty = true;
    }

    if (isCustomEmpty) {
      notify({
        failed: true,
        title: `Could not import participants.`,
        body: "Custom key/value pair can't be empty.",
      });
      return;
    }

    customData = { [customKey]: setValue(customValue) };
  }

  dispatch("import", { loading: true });
  notify({
    failed: false,
    title: `Importing participants.`,
  });
  setOpen(false);
  getParticipants(file, customData, async (newParticipants, error) => {
    if (error) {
      notify({
        failed: true,
        title: `Could not import participants.`,
        body: error,
      });
      dispatch("import", { loading: false });
      return;
    }

    try {
      await mutate(client, {
        mutation: ADD_PARTICIPANTS,
        variables: {
          input: {
            participants: newParticipants,
            projectID,
          },
        },
      });

      files = [];
      customKey = null;
      customValue = null;
      notify({
        failed: false,
        title: `Participants imported.`,
      });
      dispatch("import", { loading: false });
      setTimeout(() => {
        location.reload();
      }, 2000);
    } catch (error) {
      console.log("error", error);
      handleErrorMessage(error);
      notify({
        failed: true,
        title: `Could not import participants`,
        body:
          "Something happened on the server, and we could not import the participants.",
      });
      dispatch("import", { loading: false });
    }
  });
}
