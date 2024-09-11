const assert = require("node:assert");
const configHook = require("../../src/hooks/config");
const test = require("node:test");
const { faker } = require("@faker-js/faker");

test("config hook", async () => {
  const options = {
    config: {
      configDir: faker.system.directoryPath(),
    },
  };

  await configHook(options);

  assert.equal(global.config, options.config);
});
