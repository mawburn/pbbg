import clone from 'lodash/cloneDeep'

export enum MapActions {
  LOAD = 'map/load',
}

const initialState: GameMap.Map = {
  coords: {},
  systems: [
    {
      id: '',
      sectors: [[]],
    },
  ],
}

export default function mapReducer(
  state: GameMap.Map = { ...initialState },
  action: any
) {
  if (action.type === MapActions.LOAD) {
    const coords: GameMap.Coords = {}

    action.payload.systems.forEach((sys: GameMap.System, sysIdx: number) => {
      sys.sectors.forEach((row, rowIdx) => {
        row.forEach((sec, secIdx) => {
          coords[sec.id] = {
            system: sysIdx,
            x: secIdx,
            y: rowIdx,
          }
        })
      })
    })

    return clone({ ...action.payload, coords })
  }

  return state
}
