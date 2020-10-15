const HTMLParser = require("node-html-parser");
const fs = require("fs");

const isSandbox = true;
const EXPORTED_FILE_NAME = `../../internal/mturk/quals/${
  isSandbox ? "sandbox" : "prod"
}.json`;
const INPUT_FILE_NAME = "input.html";
const root = HTMLParser.parse(fs.readFileSync(INPUT_FILE_NAME));

const getQualsJSON = () => {
  const extractQual = (quals, typeQual) =>
    quals.childNodes.map((q) => {
      let type = typeQual === "system" ? "COMPARISON" : "BOOL";
      // if qualification type is a location
      if (q.getAttribute("value") === "00000000000000000071") {
        type = "LOCATION";
      }

      return {
        id: q.getAttribute("value"),
        name: q.childNodes[0].rawText.replace(" &amp;", " &"),
        type,
      };
    });
  const PREMIUM_QUALIFICATIONS = "Premium Qualifications";
  const SYSTEM_QUALIFICATIONS = "System Qualifications";

  const quals = root.querySelectorAll("optgroup");
  const systemQuals = quals.find(
    (q) => q.getAttribute("label") === SYSTEM_QUALIFICATIONS
  );
  const premiumQuals = quals.find(
    (q) => q.getAttribute("label") === PREMIUM_QUALIFICATIONS
  );

  return {
    qualtypes: [
      {
        type: "system",
        name: SYSTEM_QUALIFICATIONS,
        items: extractQual(systemQuals, "system"),
      },
      {
        type: "premium",
        name: PREMIUM_QUALIFICATIONS,
        items: extractQual(premiumQuals, "premium"),
      },
    ],
  };
};

fs.writeFileSync(EXPORTED_FILE_NAME, JSON.stringify(getQualsJSON(), "", 2));
