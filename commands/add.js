import fs from "fs-extra";
import path from "path";
import prompts from "prompts";
import { Command } from "@oclif/core";

export default class Add extends Command {
  static description = "save database connection";

  #questions = [
    {
      type: "text",
      name: "host",
      message: "host",
    },
    {
      type: "text",
      name: "username",
      message: "username",
    },
    {
      type: "text",
      name: "password",
      message: "password",
      style: "password",
    },
  ];

  #sqliteQuestions = [
    {
      type: "text",
      name: "storage",
      message: "path to .sqlite file",
    },
  ];

  async run() {
    const { dialect } = await prompts({
      choices: [
        {
          title: "sqlite",
          value: "sqlite",
        },
        {
          title: "mysql",
          value: "mysql",
        },
      ],
      message: "dialect",
      name: "dialect",
      type: "select",
    });

    const answers = await prompts(
      dialect == "sqlite" ? this.#sqliteQuestions : this.#questions,
    );

    const configPath = path.join(this.config.configDir, "config.json");

    await fs.ensureFile(configPath);

    const config = (await fs.readJson(configPath, { throws: false })) ?? {};

    if (!config.hasOwnProperty("databases")) {
      config.databases = {};
    }

    config.databases[answers.host || answers.storage] = {
      ...answers,
      dialect,
    };

    await fs.writeJson(configPath, config);
  }
}
