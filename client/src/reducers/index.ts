import { combineReducers } from 'redux'

import mapReducer from './map'
import sectorReducer, { Sector } from './sector'

export interface AppState {
  map: GameMap.Map
  sector: Sector
}

const rootReducer = () =>
  combineReducers<AppState>({
    map: mapReducer,
    sector: sectorReducer,
  })

export default rootReducer
