export function download(
  content,
  filename,
  mime = "application/json;charset=utf-8"
) {
  var el = document.createElement("a");
  const href = `data:${mime},${encodeURIComponent(content)}`;
  el.setAttribute("href", href);
  el.setAttribute("download", filename);
  el.style.display = "none";
  document.body.appendChild(el);
  el.click();
  document.body.removeChild(el);
}
