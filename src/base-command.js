const fs = require("fs-extra");
const path = require("path");
const { Command } = require("@oclif/core");
const History = require("./api/history");

/** @abstract */
module.exports = class BaseCommand extends Command {
  #configPath = path.join(this.config.configDir, "config.json");

  constructor(argv, config) {
    super(argv, config);

    /** @type {History} */
    this.history = new History(this.config.configDir);
  }

  async getConfig() {
    if (await fs.exists(this.#configPath)) {
      return await fs.readJson(this.#configPath, {
        throws: false,
      });
    }

    return { databases: {} };
  }

  async setConfig(config) {
    await fs.writeJson(this.#configPath, config);
  }
};
