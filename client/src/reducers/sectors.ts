import clone from 'lodash/cloneDeep'

export interface Sectors {
  [key: string]: Sector
}

export interface CurrentSector extends Sector {
  id?: string
}

export interface SectorState {
  currentSector: CurrentSector | null
  sectors: Sectors | undefined
}

export enum SectorsActionTypes {
  UPDATE = 'sectors/update_player',
  LOAD = 'sectors/load',
}

export interface SectorsAction {
  type: SectorsActionTypes
  payload?: Sectors
  playerUpdate?: {
    sectorId: string
    players: string[]
  }
}

const initialState: SectorState = {
  currentSector: null,
  sectors: {},
}

export default function sectorsReducer(
  state: SectorState = initialState,
  action: SectorsAction
) {
  if (action.type === SectorsActionTypes.LOAD) {
    return { ...state, sectors: clone(action.payload) }
  } else if (
    action.type === SectorsActionTypes.UPDATE &&
    action.playerUpdate &&
    state.sectors
  ) {
    const currentSector: CurrentSector =
      clone(state.sectors[action.playerUpdate.sectorId]) || {}

    currentSector.id = action.playerUpdate.sectorId
    currentSector.players = [...action.playerUpdate.players]

    return { ...state, currentSector }
  }

  return state
}
