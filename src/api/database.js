const fs = require("fs-extra");
const path = require("path");

/** @returns {Promise<object>} */
async function getDatabases() {
  await fs.ensureFile(getPath());

  const databases = await fs.readJson(getPath(), {
    throws: false,
  });

  return databases || {};
}

/**
 * @param {string} name
 * @param {object} database
 * @param {string} database.username
 * @param {string} database.password
 * @param {string} database.dialect
 * @param {string} [database.host]
 * @param {string} [database.storage]
 */
async function saveDatabase(name, database) {
  const databases = await getDatabases();
  databases[name] = database;
  await fs.writeJson(getPath(), databases);
}

/** @returns {string} */
function getPath() {
  return path.join(global.config.configDir, "databases.json");
}

module.exports = {
  saveDatabase,
};
