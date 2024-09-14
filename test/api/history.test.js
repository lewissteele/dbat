const mock = require("mock-fs");
const path = require("path");
const { faker } = require("@faker-js/faker");
const { getHistory, pushToHistory } = require("../../src/api/history");

test("history saves to file", async () => {
  global.config = {
    configDir: faker.system.directoryPath(),
  };

  const configFile = path.join(global.config.configDir, "history.json");

  mock({
    [configFile]: "",
  });

  const query = "select * from test";

  await pushToHistory(query);

  expect((await getHistory()).includes(query)).toBeTruthy()
});
