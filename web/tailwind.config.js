module.exports = {
  theme: {
    extend: {
      fontFamily: {
        sans: ["Open Sans", "sans-serif"],
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
  plugins: [],
};
