import clone from 'lodash/cloneDeep'

export enum MapActions {
  LOAD = 'map/load',
}
  
const initialState: GameMap.Map = {
 systems: [{
   id: '',
   sectors: [[]],
 }]
}

export default function mapReducer(state: GameMap.Map = { ...initialState }, action: any) {
  if (action.type === MapActions.LOAD) {
    return clone(action.payload)
  }

  return state
}
