const BaseCommand = require("../base-command");
const prompts = require("prompts");

export default class Add extends BaseCommand {
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
        {
          title: "mariadb",
          value: "mariadb",
        },
        {
          title: "postgres",
          value: "postgres",
        },
      ],
      message: "dialect",
      name: "dialect",
      type: "select",
    });

    const answers = await prompts(
      dialect == "sqlite" ? this.#sqliteQuestions : this.#questions,
    );

    const config = await this.getConfig();

    config.databases[answers.host || answers.storage] = {
      ...answers,
      dialect,
    };

    await this.setConfig(config);
  }
}
