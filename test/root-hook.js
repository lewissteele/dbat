const fs = require("fs-extra");

exports.mochaHooks = {
  beforeAll: () => (global.config = { configDir: "config" }),
  afterAll: () => fs.remove("config"),
};
