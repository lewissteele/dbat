const path = require("path");
const prompts = require("prompts");
const questions = require("../api/questions");
const { Command } = require("@oclif/core");
const { saveDatabase } = require("../api/database");

module.exports = class Add extends Command {
  static description = "save database connection";

  async run() {
    const database = {};

    Object.assign(database, await prompts({
      choices: questions.dialects,
      message: "dialect",
      name: "dialect",
      type: "select",
    }));

    Object.assign(database, await prompts(questions[database.dialect]));

    const { name } = await prompts({
      type: "text",
      name: "name",
      message: "name",
      initial: database.host || path.parse(database.storage).base,
    });

    await saveDatabase(name, database);
  }
};
