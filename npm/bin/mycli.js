#!/usr/bin/env node
const { spawnSync } = require("node:child_process");
const fs = require("node:fs");
const { binaryPath, install } = require("../lib/binary");

async function main() {
  let binary = binaryPath();
  if (!fs.existsSync(binary)) {
    binary = await install();
  }

  const result = spawnSync(binary, process.argv.slice(2), { stdio: "inherit" });
  if (result.error) {
    throw result.error;
  }
  process.exit(result.status ?? 1);
}

main().catch((error) => {
  console.error(error.message);
  process.exit(1);
});
