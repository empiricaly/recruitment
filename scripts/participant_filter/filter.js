var readline = require("readline");
var rl = readline.createInterface({
  input: process.stdin,
  terminal: false,
});

let js = "";
let filter;
let args = [];
let argsDone = false;
let participants = [];

rl.on("close", function () {
  args.push(participants);
  const result = filter.apply(null, args);

  console.log(result);
  result.forEach((p) => console.log(p));
});

rl.on("line", function (line) {
  if (!filter) {
    if (line.includes("END_OF_JS")) {
      filter = Function('"use strict";return (' + js + ")(")();
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
    args.push(JSON.parse(line));
  }

  participants.push(JSON.parse(line));
});
