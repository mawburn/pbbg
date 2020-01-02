import { combineReducers } from 'redux'

import sectorsReducer, { SectorState } from './sectors'
import systemsReducer, { Systems } from './systems'
import userReducer, {User} from './user'

export interface AppState {
  sectors: SectorState
  systems: Systems
  user: User
}

const rootReducer = () =>
  combineReducers<AppState>({
    sectors: sectorsReducer,
    systems: systemsReducer,
    user: userReducer
  })

export default rootReducer
