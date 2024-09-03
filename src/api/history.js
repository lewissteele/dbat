const fs = require("fs-extra");
const path = require("path");

/** @param {string} query */
module.exports.push = async function (query) {
  const history = await getHistory();
  history.push(query);
  await fs.writeJson(getPath(), history);
};

/** @returns {string} */
function getPath() {
  return path.join(global.config.configDir, "history.json");
}

/** @returns {array<string>} */
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
