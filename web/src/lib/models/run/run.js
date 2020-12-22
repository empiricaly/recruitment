import dayjs from "dayjs";

import { download } from "../../../utils/download";

export function exportJson(template) {
  const content = JSON.stringify(template);
  const mime = "application/json;charset=utf-8";
  const date = dayjs().format("YYYY-MM-DDTHH:mm:ss");
  const filename = `Empirica recruitment export Run - ${template.name} – ${date}.json`;
  download(content, filename, mime);
}
