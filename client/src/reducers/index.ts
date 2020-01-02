import { combineReducers } from 'redux'

import sectorsReducer, { SectorState } from './sectors'
import systemsReducer, { Systems } from './systems'

export interface AppState {
  sectors: SectorState
  systems: Systems
}

const rootReducer = () =>
  combineReducers<AppState>({
    sectors: sectorsReducer,
    systems: systemsReducer,
  })

export default rootReducer
