const assert = require("node:assert");
const { getHistory, pushToHistory } = require("../../src/api/history");

it("saves history to file", async () => {
  const query = "select * from test";

  await pushToHistory(query);

  const history = await getHistory();

  assert(history.includes(query));
});
