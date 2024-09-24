const List = require("../../src/commands/list");
const assert = require("node:assert");
const { captureOutput } = require("@oclif/test");
const { saveDatabase, removeDatabase } = require("../../src/api/database");

beforeEach(() => removeDatabase("test"));

it("lists all saved database connections", async () => {
  await saveDatabase("test", {
    dialect: "sqlite",
    storage: ":memory:",
  });

  const { stdout } = await captureOutput(() => List.run());

  assert.equal(stdout.trim(), "test");
});

it("still works when there are no databases", async () => {
  const { stdout } = await captureOutput(() => List.run());

  assert.equal(stdout.trim(), "no databases");
});
