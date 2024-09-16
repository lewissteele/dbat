const assert = require("node:assert");
const configHook = require("../../src/hooks/config");

it("sets config to global object", async () => {
  const options = {
    config: { configDir: "config" },
  };

  await configHook(options);

  assert.equal(global.config, options.config);
});
