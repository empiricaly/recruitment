const HTMLParser = require("node-html-parser");
const fs = require("fs");

const EXPORTED_FILE_NAME = "quals.json";
const INPUT_FILE_NAME = "input.html";
const root = HTMLParser.parse(fs.readFileSync(INPUT_FILE_NAME));

const getQualsJSON = () => {
  const extractQual = (quals) =>
    quals.childNodes.map((q) => {
      return {
        id: q.getAttribute("value"),
        name: q.childNodes[0].rawText.replace(" &amp;", " &"),
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
        items: extractQual(systemQuals),
      },
      {
        type: "premium",
        name: PREMIUM_QUALIFICATIONS,
        items: extractQual(premiumQuals),
      },
    ],
  };
};

fs.writeFileSync(EXPORTED_FILE_NAME, JSON.stringify(getQualsJSON(), "", 2));
