import streamSaver from "streamsaver";

export function download(
  content,
  filename,
  mime = "application/json;charset=utf-8"
) {
  const uInt8 = new TextEncoder().encode(content);
  const fileStream = streamSaver.createWriteStream(filename, {
    size: uInt8.byteLength,
  });

  const writer = fileStream.getWriter();
  writer.write(uInt8);
  writer.close();
}
