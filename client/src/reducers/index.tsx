import { combineReducers } from 'redux'

import mapReducer from './map'

export interface AppState {
  map: any
}

const rootReducer = () =>
  combineReducers<AppState>({
    map: mapReducer,
  })

export default rootReducer
