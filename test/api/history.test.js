const { getHistory, pushToHistory } = require("../../src/api/history");

test("history saves to file", async () => {
  const query = "select * from test";

  await pushToHistory(query);

  const history = await getHistory();

  expect(history.includes(query)).toBe(true);
});
