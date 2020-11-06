// echo 'function (participants, arg1, arg2) { return participants.reverse() }\n\n\nEND_OF_JS\n{"yup": 1}\n{"huh": 2}\nEND_OF_ARGS\n{"id": "123"}\n{"id": "321"}' | node scripts/participant_filter/filter.js

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
  for (let i = 0; i < participants.length; i++) {
    participants[i].get = (k) => {
      if (participants[i].changes && participants[i].changes[k] !== undefined) {
        return participants[i].changes[k];
      }

      const data = participants[i].edges.Data;
      if (!data) {
        return;
      }

      const datum = data.find((d) => d.current && d.key === k);
      if (datum) {
        return JSON.parse(datum.val);
      }
    };
    participants[i].set = (k, v) => {
      if (!participants[i].changes) {
        participants[i].changes = {};
      }

      participants[i].changes[k] = v;
    };
  }
  args.unshift(participants);
  const result = filter.apply(null, args);
  const keeps = {};
  result.forEach((p) => {
    keeps[p.id] = true;
  });

  participants.forEach((p) => {
    process.stdout.write(
      JSON.stringify({
        id: p.id,
        keep: Boolean(keeps[p.id]),
        changes: p.changes,
      }) + "\n"
    );
  });
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

  participants.push(JSON.parse(line));
});
