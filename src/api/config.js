const fs = require("fs-extra");
const path = require("path");

module.exports.getConfig = async function () {
  const path = getPath();

  if (await fs.exists(path)) {
    return await fs.readJson(path, {
      throws: false,
    });
  }

  return { databases: {} };
};

module.exports.setConfig = async function (config) {
  await fs.writeJson(getPath(), config);
};

function getPath() {
  return path.join(global.config.configDir, "config.json");
}
