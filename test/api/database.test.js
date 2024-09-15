const assert = require("node:assert");
const fs = require("fs-extra");
const { test, before, after } = require("node:test");
const {
  saveDatabase,
  getDatabases,
  getConnection,
} = require("../../src/api/database");

before(() => {
  global.config = { configDir: "config" };
});

test("database connection is saved to file", async () => {
  const database = {
    dialect: "sqlite",
    storage: ":memory:",
  };

  await saveDatabase("test", database);

  const databases = await getDatabases();

  assert(databases["test"], database);
});

test("it can connect to database", async () => {
  await saveDatabase("test", {
    dialect: "sqlite",
    storage: ":memory:",
  });

  const connection = await getConnection("test");

  connection.authenticate();
});

after(async () => {
  await fs.rm(global.config.configDir, { recursive: true });
});
