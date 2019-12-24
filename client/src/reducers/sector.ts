import clone from 'lodash/cloneDeep'

export interface Sector {
  x: number
  y: number
}

export enum SectorActions {
  MOVE = 'sector/move',
}

const initialState = {
  x: 0,
  y: 0,
}

export default function sectorReducer(
  state: object | null = { ...initialState },
  action: any
) {
  if (action.type === SectorActions.MOVE) {
    return clone(action.payload)
  }

  return state
}
