const List = require("../../src/commands/list");
const assert = require("node:assert");
const fs = require("fs-extra");
const { captureOutput } = require("@oclif/test");
const { saveDatabase, removeDatabase } = require("../../src/api/database");

it("lists all saved database connections", async () => {
  const expected = "test";

  saveDatabase(expected, {
    dialect: "sqlite",
    storage: ":memory:",
  });

  const { stdout } = await captureOutput(() => List.run());

  assert.equal(stdout.trim(), expected);

  removeDatabase(removeDatabase);
});

it("still works when there are no databases", async () => {
  const { stdout } = await captureOutput(() => List.run());

  assert.equal(stdout.trim(), "no databases");
});
