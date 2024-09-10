const { getConfig } = require("../api/config");
const { Command } = require("@oclif/core");

module.exports = class List extends Command {
  static description = "show databases";
  async run() {
    const config = await getConfig();
    const databases = Object.keys(config.databases);

    if (!databases.length) {
      this.log("no databases");
      return;
    }

    databases.forEach((val) => this.log(val));
  }
};
