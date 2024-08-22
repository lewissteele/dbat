import fs from "fs-extra";
import path from "path";
import { Command } from "@oclif/core";

/** @abstract */
export default class BaseCommand extends Command {
  #defaultConfig = {
    databases: [],
  };

  async getConfig() {
    await fs.ensureFile(this.#getConfigPath())

    const config = await fs.readJson(this.#getConfigPath(), {
      throws: false,
    });

    if (config) {
      return config;
    }

    return this.#defaultConfig;
  }

  async setConfig(config) {
    await fs.writeJson(this.#getConfigPath(), config);
  }

  #getConfigPath() {
    return path.join(this.config.configDir, "config.json");
  }
}
