const BaseCommand = require("../base-command");
const Table = require("cli-table3");
const readline = require("node:readline/promises");
const { Args } = require("@oclif/core");
const { Sequelize } = require("sequelize");
const { stdin, stdout } = require("node:process");

export default class Connect extends BaseCommand {
  static args = {
    database: Args.string(),
  };

  async run() {
    const { args } = await this.parse(Connect);

    const config = await this.getConfig();
    const database = config.databases[args.database];

    if (database == undefined) {
      this.log("no such database");
      return;
    }

    const sequelize = new Sequelize({
      ...database,
      logging: false,
    });

    const rl = readline.createInterface({
      input: stdin,
      output: stdout,
    });

    rl.prompt();

    rl.on("line", async (line) => {
      const [results, meta] = await sequelize.query(line, { raw: true });

      if (!results.length) {
        rl.prompt();
        return;
      }

      const headers = Object.keys(results[0]);
      const rows = results.map((row) => Object.values(row));

      const table = new Table({
        head: headers,
        rows: rows,
      });

      table.push(...rows);

      this.log(table.toString());

      rl.prompt();
    });
  }
}
