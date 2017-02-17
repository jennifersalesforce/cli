const fs = require('fs-extra')
const path = require('path')
let version
try {
  version = fs.readFileSync(path.join(__dirname, '..', 'VERSION'), 'utf-8').trim()
} catch (err) {
  version = require('../package.json').version
}

module.exports = version