const fs = require("fs-extra");
const path = require("path");

/** @param {string} configDir */
function History(configDir) {
  this.historyPath = path.join(configDir, "history.json");
}

/** @returns {array<string>} */
History.prototype.getHistory = async function () {
  await fs.ensureFile(this.historyPath);

  const history = await fs.readJson(this.historyPath, {
    throws: false,
  });

  if (Array.isArray(history)) {
    return history;
  }

  return [];
};

/** @param {string} query */
History.prototype.push = async function (query) {
  const history = await this.getHistory();
  history.push(query);
  await fs.writeJson(this.historyPath, history);
};

module.exports = History;
