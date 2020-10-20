const tailwindcss = require("tailwindcss");
const cssnano = require("cssnano");
const postcssImport = require("postcss-import");
const presetEnv = require("postcss-preset-env")({
  features: {
    // enable nesting
    "nesting-rules": true,
  },
});

// only needed if you want to purge
const purgecss = require("@fullhuman/postcss-purgecss")({
  content: ["./src/**/*.svelte", "./public/**/*.html"],
  whitelistPatterns: [
    /svelte-/,
    // Needed for callout and badges
    /^txt-[a-z]+?-\d{2,3}$/,
    /^bg-[a-z]+?-\d{2,3}$/,
  ],
  defaultExtractor: (content) => content.match(/[\w-/.:]+(?<!:)/g) || [],
});

module.exports = {
  plugins: [
    postcssImport,
    tailwindcss("./tailwind.config.js"),
    // only needed if you want to purge
    ...(process.env.NODE_ENV === "production"
      ? [purgecss, presetEnv, cssnano]
      : []),
  ],
};
