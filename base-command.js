const fs = require("fs-extra");
const path = require("path");
const { Command } = require("@oclif/core");

/** @abstract */
module.exports = class BaseCommand extends Command {
  #configPath = path.join(this.config.configDir, "config.json");

  #defaultConfig = { databases: [] };

  async getConfig() {
    await fs.ensureFile(this.#configPath);

    const config = await fs.readJson(this.#configPath, {
      throws: false,
    });

    if (config) {
      return config;
    }

    return this.#defaultConfig;
  }

  async setConfig(config) {
    await fs.writeJson(this.#configPath, config);
  }
}
