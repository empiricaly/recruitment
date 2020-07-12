const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  theme: {
    extend: {
      fontFamily: {
        // sans: ["Open Sans", "sans-serif"],
        sans: ["Inter var", ...defaultTheme.fontFamily.sans],
      },

      width: {
        "128": "32rem",
      },

      boxShadow: {
        outline: "0 0 0 3px #63B3ED",
      },
    },
  },
  variants: {
    boxShadow: ["responsive", "hover", "focus", "active", "group-hover"],
    borderColor: ["responsive", "hover", "focus", "active", "group-hover"],
    backgroundColor: ["responsive", "hover", "focus", "active", "group-hover"],
  },
  plugins: [
    require("@tailwindcss/ui")({
      layout: "sidebar",
    }),
  ],
};
