const configHook = require("../../src/hooks/config");
const { faker } = require("@faker-js/faker");

test("config hook", async () => {
  const options = {
    config: {
      configDir: faker.system.directoryPath(),
    },
  };

  await configHook(options);

  expect(global.config).toEqual(options.config);
});
