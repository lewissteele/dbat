const fs = require("fs-extra");
const path = require("path");

/** @returns {Promise<array<string>>} */
async function getHistory() {
  await fs.ensureFile(getPath());

  const history = await fs.readJson(getPath(), {
    throws: false,
  });

  if (Array.isArray(history)) {
    return history;
  }

  return [];
}

/** @param {string} query */
async function pushToHistory(query) {
  const history = await getHistory();
  history.push(query);
  await fs.writeJson(getPath(), history);
}

/** @returns {string} */
function getPath() {
  return path.join(global.config.configDir, "history.json");
}

module.exports = {
  getHistory,
  pushToHistory,
};
