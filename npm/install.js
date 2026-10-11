const { install } = require("./lib/binary");

install().catch((error) => {
  console.error(error.message);
  process.exit(1);
});
