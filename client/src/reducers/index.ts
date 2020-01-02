import { combineReducers } from 'redux'

import playerSectorReducer, { PlayerSector } from './playerSector'
import sectorsReducer, { Sectors } from './sectors'
import systemsReducer, { Systems } from './systems'

export interface AppState {
  playerSector: PlayerSector
  sectors: Sectors
  systems: Systems
}

const rootReducer = () =>
  combineReducers<AppState>({
    playerSector: playerSectorReducer,
    sectors: sectorsReducer,
    systems: systemsReducer,
  })

export default rootReducer
