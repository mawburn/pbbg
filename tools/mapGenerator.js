const fs = require('fs')

const ROW_COUNT = 4
const COL_COUNT = 4

function randomNum(max) {
  return Math.floor(Math.random() * (max - 1) + 1)
}

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

  for (let i = 0; i < 12; i++) {
    const _shouldPopulate = randomNum(25)
    let _object = null

    if (_shouldPopulate === 1) {
      const sizeRan = randomNum(4)
      
      _object = {
        id: randomString(7),
        size: sizeRan === 2 ? 'lg' : sizeRan === 1 ? 'md' : 'sm',
        type: 'ore',
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
    const _celRan = randomNum(10)
    const celestial = _celRan === 1 ? {
      id: randomString(7),
      name: randomString(3),
      type: 'station'
    } : _celRan === 2 ? {
      id: randomString(7),
      name: randomString(3),
      type: 'planet'
    } : null

    const sector = {
      id: randomString(7),
      celestial,
      objects: sectorObjCreator(),
    }

    sectorRow.push(sector)
  }

  system.sectors.push(sectorRow)
}

fs.writeFileSync(
  `./output/system-${system.id}.json`,
  JSON.stringify({
    systems: [system]
  }, null, 2),
  err => {
    if (err) throw new Error(err)
  }
)