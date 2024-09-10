const fs = require("fs-extra");
const path = require("path");

async function getConfig() {
  const path = getPath();

  if (await fs.exists(path)) {
    return await fs.readJson(path, {
      throws: false,
    });
  }

  return { databases: {} };
}

async function setConfig(config) {
  await fs.writeJson(getPath(), config);
}

function getPath() {
  return path.join(global.config.configDir, "config.json");
}

module.exports = {
  getConfig,
  setConfig,
};
