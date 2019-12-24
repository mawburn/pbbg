import clone from 'lodash/cloneDeep'

export enum MapActions {
  LOAD = 'map/load',
}

export default function mapReducer(state: object | null = null, action: any) {
  if (action.type === MapActions.LOAD) {
    return clone(action.payload)
  }

  return state
}
