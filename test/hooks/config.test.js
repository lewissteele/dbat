const assert = require("node:assert");
const test = require("node:test");
const configHook = require("../../src/hooks/config");

test("config hook", async () => {
  const options = {
    config: {
      configDir: "/home/lewis/.config/dbat",
    },
  };

  await configHook(options);

  assert.equal(global.config, options.config);
});
