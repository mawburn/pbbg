import clone from 'lodash/cloneDeep'

export interface PlayerSector {
  id: string
  players?: string[]
}

export enum PlayerSectorActionTypes {
  LOAD = 'playerSector/load',
}

export interface PlayerSectorAction {
  type: PlayerSectorActionTypes
  payload: PlayerSector
}

const initialState: PlayerSector = {
  id: '',
  players: []
}

export default function playerSectorReducer(state: PlayerSector = clone(initialState), action: PlayerSectorAction) {
  if (action.type === PlayerSectorActionTypes.LOAD) {
    return clone(action.payload)
  }

  return state
}
