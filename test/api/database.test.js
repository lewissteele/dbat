const fs = require("fs-extra");
const {
  saveDatabase,
  getDatabases,
  getConnection,
} = require("../../src/api/database");

const configDir = "./config";

beforeAll(() => {
  global.config = { configDir: configDir };
});

test("database connection is saved to file", async () => {
  const database = {
    dialect: "sqlite",
    storage: ":memory:",
  };

  await saveDatabase("test", database);

  const databases = await getDatabases();

  expect(databases["test"]).toEqual(database);
});

test("it can connect to database", async () => {
  jest.mock('fs-extra')

  await saveDatabase("test", {
    dialect: "sqlite",
    storage: ":memory:",
  });

  const connection = await getConnection("test");

  connection.authenticate();
});

afterAll(async () => {
  await fs.rm(configDir, { recursive: true });
});
