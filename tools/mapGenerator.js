const fs = require('fs')

const ROW_COUNT = 4
const COL_COUNT = 4

function randomString(length) {
  let result = ''

  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'

  for (let i = 0; i < length; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }

  return result
}

function sectorObjCreator() {
  const _objects = []

  for (let i = 0; i < 16; i++) {
    const _shouldPopulate = Math.floor(Math.random() * (42 - 1) + 1)
    let _object = null

    if (_shouldPopulate === 1) {
      _object = {
        id: randomString(7),
        type: 'ore',
        subType: 'crystal',
        quantity: Math.floor(Math.random() * (1000000 - 100) + 100),
      }
    }

    _objects.push(_object)
  }

  return _objects
}

const system = {}

system.id = randomString(7)
system.sectors = []

for (let row = 0; row < ROW_COUNT; row++) {
  const sectorRow = []

  for (let col = 0; col < COL_COUNT; col++) {
    const sector = {
      id: randomString(7),
      objects: sectorObjCreator(),
    }

    sectorRow.push(sector)
  }

  system.sectors.push(sectorRow)
}

fs.writeFileSync(
  `./output/system-${system.id}.json`,
  JSON.stringify(system, null, 2),
  err => {
    if (err) throw new Error(err)
  }
)
