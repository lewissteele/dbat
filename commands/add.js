import { Command } from "@oclif/core";
import prompts from "prompts";

export default class Add extends Command {
  static description = "save database connection";
  static aliases = ["save"];

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
    },
  ];

  async run() {
    const answers = await prompts(this.#questions);

    this.log(answers);
  }
}
