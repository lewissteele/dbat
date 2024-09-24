const fs = require("fs-extra");
const path = require("path");
const { Sequelize } = require("sequelize");

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

async function removeDatabase(name) {
  const databases = await getDatabases();
  delete databases[name];
  await fs.writeJson(getPath(), databases);
}

/**
 * @param {string} name
 * @returns {Promise<Sequelize|null>}
 */
async function getConnection(name) {
  const database = (await getDatabases())[name];

  if (!database) {
    return null;
  }

  return new Sequelize({
    ...database,
    logging: false,
  });
}

/** @returns {string} */
function getPath() {
  return path.join(global.config.configDir, "databases.json");
}

module.exports = {
  getConnection,
  getDatabases,
  removeDatabase,
  saveDatabase,
};
