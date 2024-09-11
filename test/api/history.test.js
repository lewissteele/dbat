const history = require("../../src/api/history");
const mock = require("mock-fs");
const path = require("path");
const test = require("node:test");
const { faker } = require("@faker-js/faker");

test("history saves to file", async () => {
  global.config = {
    configDir: faker.system.directoryPath(),
  };

  const configFile = path.join(
    global.config.configDir,
    "history.json",
  );

  mock({
    [configFile]: '',
  });

  await history.push("select * from test");
});
