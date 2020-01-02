const fs = require('fs')

const shortid = require('shortid')
const rWords = require('random-words')

const ROW_COUNT = 7
const COL_COUNT = 10

const randomNum = (min, max) => Math.floor(Math.random() * (max - min) + min)

const getCelestial = () => {
  const _r = randomNum(1, 7)

  if (_r !== 1) {
    return null
  }

  return {
    id: shortid.generate(),
    name: rWords(2).join(' '),
    type: randomNum(1, 3) === 1 ? 'station' : 'planet',
  }
}

const getObjects = () => {
  const objects = []

  for (let i = 0; i < 6; i++) {
    if (randomNum(1, 10) === 1) {
      const max = Math.round(randomNum(10000, 10000000) / 10000) * 10000

      objects.push({
        id: shortid.generate(),
        max,
        type: 'ore',
      })
    } else {
      objects.push(null)
    }
  }

  return objects
}

const getSectors = () => {
  const total = ROW_COUNT * COL_COUNT
  const sectors = {}

  for (let i = 0; i < total; i++) {
    const id = i.toString(36).padStart(4, '0')

    sectors[id] = {
      celestial: getCelestial(),
      objects: [...getObjects()],
    }
  }

  return sectors
}

function main() {
  const systemId = 'S' + (1).toString(36).padStart(3, '0')
  const sectors = getSectors()
  const sectorIds = Object.keys(sectors)

  const map = {
    systems: {
      [systemId]: [],
    },
    sectors,
  }

  for (let row = 0; row < ROW_COUNT; row++) {
    const rowArr = []

    for (let col = 0; col < COL_COUNT; col++) {
      const id = sectorIds.shift()

      rowArr.push(id)

      sectors[id].x = col
      sectors[id].y = row
      sectors[id].system = systemId
    }

    map.systems[systemId].push(rowArr)
  }

  fs.writeFileSync(
    `./output/map.json`,
    JSON.stringify(map, null, 2),
    { encoding: 'utf8', flag: 'w' },
    err => {
      if (err) throw new Error(err)
    }
  )
}

main()
