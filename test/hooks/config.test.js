const configHook = require("../../src/hooks/config");

test("config hook", async () => {
  const options = {
    config: { dir: "config" },
  };

  await configHook(options);

  expect(global.config).toEqual(options.config);
});
