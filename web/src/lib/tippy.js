import tippy from "tippy.js";

export default function (node, props) {
  tippy(node, { delay: [0, 100], theme: "translucent", ...props });
}
