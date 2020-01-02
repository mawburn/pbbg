import clone from 'lodash/cloneDeep'

export interface Systems {
  [key: string]: System
}

export enum SystemsActionTypes {
  LOAD = 'systems/load',
}

export interface SystemsAction {
  type: SystemsActionTypes
  payload: Systems
}

export default function sectorsReducer(state: Systems = {}, action: SystemsAction) {
  if (action.type === SystemsActionTypes.LOAD) {
    return clone(action.payload)
  }

  return state
}
