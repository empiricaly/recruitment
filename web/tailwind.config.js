const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  theme: {
    extend: {
      fontFamily: {
        // sans: ["Open Sans", "sans-serif"],
        sans: ["Inter var", ...defaultTheme.fontFamily.sans],
      },

      width: {
        128: "32rem",
      },

      boxShadow: {
        outline: "0 0 0 3px #63B3ED",
      },

      colors: {
        mint: {
          900: "#1C403C",
          800: "#265550",
          700: "#32716B",
          600: "#3F8D85",
          500: "#4EB1A7",
          400: "#80C7BF",
          300: "#A5D7D2",
          200: "#CEE9E6",
          100: "#EDF7F6",
        },
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
