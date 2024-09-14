const Table = require("cli-table3");
const readline = require("node:readline/promises");
const { Args } = require("@oclif/core");
const { Command } = require("@oclif/core");
const { getConnection, databaseExists } = require("../api/database");
const { pushToHistory } = require("../api/history");
const { stdin, stdout } = require("node:process");

module.exports = class Connect extends Command {
  static description = "run queries on database";
  static args = {
    database: Args.string(),
  };

  async run() {
    const { args } = await this.parse(Connect);

    if (!databaseExists(args.database)) {
      this.log("no such database");
      return;
    }

    const connection = await getConnection(args.database);

    const rl = readline.createInterface({
      input: stdin,
      output: stdout,
    });

    rl.prompt();

    rl.on("line", async (query) => {
      const [results, meta] = await connection.query(query, { raw: true });

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

      await pushToHistory(query);

      rl.prompt();
    });
  }
};
