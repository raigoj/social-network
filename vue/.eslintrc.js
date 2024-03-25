module.exports = {
  root: true,

  env: {
    node: true,
  },

  extends: [
    "plugin:vue/vue3-essential",
    "eslint:recommended",
    "plugin:prettier/recommended",
  ],

  parserOptions: {},

  rules: {
    "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
  },

  // eslint-disable-next-line no-dupe-keys
  extends: [
    "plugin:vue/vue3-essential",
    "eslint:recommended",
    "plugin:prettier/recommended",
    //"@vue/typescript",
    "@vue/typescript/recommended",
  ],
};
