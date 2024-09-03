const fs = require("fs-extra");
const path = require("path");
const { Command } = require("@oclif/core");

/** @abstract */
module.exports = class BaseCommand extends Command {
  #configPath = path.join(this.config.configDir, "config.json");

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
