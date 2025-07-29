#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const https = require('https');
const { execSync } = require('child_process');

const packageJson = require('../package.json');
const version = packageJson.version;

// Map Node.js platform and arch to GoReleaser naming convention
function getPlatformInfo() {
  const platform = process.platform;
  const arch = process.arch;
  
  let goos, goarch;
  
  switch (platform) {
    case 'darwin':
      goos = 'Darwin';
      break;
    case 'linux':
      goos = 'Linux';
      break;
    case 'win32':
      goos = 'Windows';
      break;
    default:
      throw new Error(`Unsupported platform: ${platform}`);
  }
  
  switch (arch) {
    case 'x64':
      goarch = 'x86_64';
      break;
    case 'arm64':
      goarch = 'arm64';
      break;
    case 'arm':
      goarch = 'armv6';
      break;
    default:
      throw new Error(`Unsupported architecture: ${arch}`);
  }
  
  return { goos, goarch };
}

function downloadFile(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    
    https.get(url, (response) => {
      if (response.statusCode === 302 || response.statusCode === 301) {
        // Handle redirect
        return downloadFile(response.headers.location, dest);
      }
      
      if (response.statusCode !== 200) {
        reject(new Error(`Failed to download: ${response.statusCode} ${response.statusMessage}`));
        return;
      }
      
      response.pipe(file);
      
      file.on('finish', () => {
        file.close();
        resolve();
      });
      
      file.on('error', (err) => {
        fs.unlink(dest, () => {}); // Delete the file on error
        reject(err);
      });
    }).on('error', (err) => {
      reject(err);
    });
  });
}

async function install() {
  try {
    const { goos, goarch } = getPlatformInfo();
    const binaryName = goos === 'Windows' ? 'presidium-oapi3.exe' : 'presidium-oapi3';
    const archiveName = `presidium-oapi3_${version}_${goos}_${goarch}.tar.gz`;
    const downloadUrl = `https://github.com/SPANDigital/presidium-oapi3/releases/download/v${version}/${archiveName}`;
    
    console.log(`Installing presidium-oapi3 v${version} for ${goos}/${goarch}...`);
    console.log(`Downloading from: ${downloadUrl}`);
    
    // Create bin directory
    const binDir = path.join(__dirname, '..', 'bin');
    if (!fs.existsSync(binDir)) {
      fs.mkdirSync(binDir, { recursive: true });
    }
    
    // Download archive
    const archivePath = path.join(binDir, archiveName);
    await downloadFile(downloadUrl, archivePath);
    
    // Extract binary
    console.log('Extracting binary...');
    if (goos === 'Windows') {
      // For Windows, we might need different extraction logic
      throw new Error('Windows extraction not yet implemented');
    } else {
      // Unix-like systems
      execSync(`tar -xzf "${archivePath}" -C "${binDir}" "${binaryName}"`, { stdio: 'inherit' });
    }
    
    // Make binary executable
    const binaryPath = path.join(binDir, binaryName);
    fs.chmodSync(binaryPath, '755');
    
    // Clean up archive
    fs.unlinkSync(archivePath);
    
    console.log(`Successfully installed presidium-oapi3 to ${binaryPath}`);
  } catch (error) {
    console.error('Installation failed:', error.message);
    process.exit(1);
  }
}

install();
