const fs = require("fs-extra");

afterAll(() => {
  if (!fs.existsSync(global.config.configDir)) {
    return;
  }

  fs.rmSync(global.config.configDir, { recursive: true });
});
