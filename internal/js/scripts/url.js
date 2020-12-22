var readline = require("readline");
var rl = readline.createInterface({
  input: process.stdin,
  terminal: false,
});

let js = "";
let filter;
let args = [];
let argsDone = false;
let steps = [];

rl.on("close", function () {
  // If participant exist on the db
  if (args[0]) {
    args[0].get = (k) => {
      const data = args[0].edges.Data;
      if (!data) {
        return;
      }

      const datum = data.find((d) => d.current && d.key === k);
      if (datum) {
        return JSON.parse(datum.val);
      }
    };
  }

  args.push(steps);
  const result = filter.apply(null, args);
  process.stdout.write(result);
});

rl.on("line", function (line) {
  process.stderr.write(`NEW LINE: ${line}\n`);
  if (!filter) {
    if (line.includes("END_OF_JS")) {
      filter = Function('"use strict";return (' + js + ")")();
      return;
    }
    js += line + "\n";
    return;
  }

  if (!argsDone) {
    if (line.includes("END_OF_ARGS")) {
      argsDone = true;
      return;
    }
    if (line.trim() !== "") {
      args.push(JSON.parse(line));
    }
    return;
  }

  steps.push(JSON.parse(line));
});
