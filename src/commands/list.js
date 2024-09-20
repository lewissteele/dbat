const _ = require("lodash");
const cmd = require("@oclif/core").Command;
const { getDatabases } = require("../api/database");

module.exports = _.create(cmd.prototype, {
  description: "show databases",
  async run() {
    const databases = Object.keys(await getDatabases());

    if (!databases.length) {
      this.log("no databases");
      return;
    }

    databases.forEach((val) => this.log(val));
  },
});
