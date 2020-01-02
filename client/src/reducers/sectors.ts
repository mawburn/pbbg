import clone from 'lodash/cloneDeep'

export interface Sectors {
  [key: string]: Sector | undefined
}

export enum SectorsActionTypes {
  LOAD = 'sectors/load',
}

export interface SectorsAction {
  type: SectorsActionTypes
  payload: Sectors
}

export default function sectorsReducer(state: Sectors = {}, action: SectorsAction) {
  if (action.type === SectorsActionTypes.LOAD) {
    return clone(action.payload)
  }

  return state
}
