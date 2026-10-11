const crypto = require("node:crypto");
const fs = require("node:fs");
const https = require("node:https");
const path = require("node:path");
const { version } = require("../package.json");

const REPOSITORY = "ArthurDescourvieres/plateforme-mycli";
const PLATFORMS = { linux: "linux", darwin: "darwin", win32: "windows" };
const ARCHITECTURES = { x64: "amd64", arm64: "arm64" };

function assetName() {
  const os = PLATFORMS[process.platform];
  const arch = ARCHITECTURES[process.arch];
  if (!os || !arch) {
    throw new Error(`mycli: unsupported platform ${process.platform}/${process.arch}`);
  }
  return `mycli_${os}_${arch}${os === "windows" ? ".exe" : ""}`;
}

function binaryPath() {
  const name = process.platform === "win32" ? "mycli.exe" : "mycli";
  return path.join(__dirname, "..", "vendor", name);
}

function download(url, redirects = 5) {
  return new Promise((resolve, reject) => {
    https
      .get(url, { headers: { "User-Agent": "mycli-npm" } }, (response) => {
        const { statusCode, headers } = response;
        if (statusCode >= 300 && statusCode < 400 && headers.location && redirects > 0) {
          response.resume();
          resolve(download(headers.location, redirects - 1));
          return;
        }
        if (statusCode !== 200) {
          response.resume();
          reject(new Error(`mycli: download failed with HTTP ${statusCode}: ${url}`));
          return;
        }
        const chunks = [];
        response.on("data", (chunk) => chunks.push(chunk));
        response.on("end", () => resolve(Buffer.concat(chunks)));
        response.on("error", reject);
      })
      .on("error", reject);
  });
}

function expectedChecksum(checksums, asset) {
  for (const line of checksums.split("\n")) {
    const [sum, name] = line.trim().split(/\s+/);
    if (name === asset) {
      return sum;
    }
  }
  throw new Error(`mycli: ${asset} is missing from checksums.txt`);
}

async function install() {
  const asset = assetName();
  const baseUrl = `https://github.com/${REPOSITORY}/releases/download/v${version}`;
  const [binary, checksums] = await Promise.all([
    download(`${baseUrl}/${asset}`),
    download(`${baseUrl}/checksums.txt`),
  ]);

  const expected = expectedChecksum(checksums.toString("utf8"), asset);
  const actual = crypto.createHash("sha256").update(binary).digest("hex");
  if (actual !== expected) {
    throw new Error(`mycli: checksum mismatch for ${asset}`);
  }

  const target = binaryPath();
  fs.mkdirSync(path.dirname(target), { recursive: true });
  fs.writeFileSync(target, binary, { mode: 0o755 });
  return target;
}

module.exports = { binaryPath, install };
