const generic = [
  {
    type: "text",
    name: "host",
    message: "host",
  },
  {
    type: "text",
    name: "username",
    message: "username",
  },
  {
    type: "text",
    name: "password",
    message: "password",
    style: "password",
  },
];

const sqlite = [
  {
    type: "text",
    name: "storage",
    message: "path to .sqlite file",
  },
];

const dialects = [
  {
    title: "sqlite",
    value: "sqlite",
  },
  {
    title: "mysql",
    value: "mysql",
  },
  {
    title: "mariadb",
    value: "mariadb",
  },
  {
    title: "postgres",
    value: "postgres",
  },
];

module.exports = {
  dialects,
  mariadb: generic,
  mysql: generic,
  postgres: generic,
  sqlite,
};
