const _ = require("lodash");
const cmd = require("@oclif/core").Command;
const { getDatabases } = require("../api/database");

module.exports = _.create(cmd.prototype, {
  description: "show databases",
  async run() {
    const databases = _.keys(await getDatabases());

    if (_.isEmpty(databases)) {
      this.log("no databases");
      return;
    }

    _.each(databases, (val) => this.log(val));
  },
});
