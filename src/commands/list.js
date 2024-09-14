const { Command } = require("@oclif/core");
const { getDatabases } = require("../api/database");

module.exports = class List extends Command {
  static description = "show databases";

  async run() {
    const databases = Object.keys(await getDatabases());

    if (!databases.length) {
      this.log("no databases");
      return;
    }

    databases.forEach((val) => this.log(val));
  }
};
