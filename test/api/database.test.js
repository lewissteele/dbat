const {
  getConnection,
  getDatabases,
  saveDatabase,
} = require("../../src/api/database");

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
  await saveDatabase("test", {
    dialect: "sqlite",
    storage: ":memory:",
  });

  const connection = await getConnection("test");

  connection.authenticate();
});
