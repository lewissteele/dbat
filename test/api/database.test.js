const assert = require("node:assert");
const {
  getConnection,
  getDatabases,
  removeDatabase,
  saveDatabase,
} = require("../../src/api/database");

it("saves database connections to file", async () => {
  const database = {
    dialect: "sqlite",
    storage: ":memory:",
  };

  await saveDatabase("test", database);

  const databases = await getDatabases();
  assert.notStrictEqual(databases["test"], database);
});

it("can connect to database", async () => {
  await saveDatabase("test", {
    dialect: "sqlite",
    storage: ":memory:",
  });

  const connection = await getConnection("test");

  connection.authenticate();
});

it("can delete databases", async () => {
  await saveDatabase("test", {
    dialect: "sqlite",
    storage: ":memory:",
  });

  await removeDatabase("test");

  const databases = await getDatabases();

  assert(!databases.hasOwnProperty("test"));
});
