const mock = require("mock-fs");
const path = require("path");
const { faker } = require("@faker-js/faker");
const { getHistory, pushToHistory } = require("../../src/api/history");

beforeEach(() => {
  const fakePath = faker.system.directoryPath();

  global.config = { configDir: fakePath };

  mock({ [path.join(fakePath, "history.json")]: "" });
});

test("history saves to file", async () => {
  const query = "select * from test";

  await pushToHistory(query);

  expect((await getHistory()).includes(query)).toBe(true);
});

afterEach(() => {
  mock.restore();
});
