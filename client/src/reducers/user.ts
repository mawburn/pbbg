import clone from 'lodash/cloneDeep'

export enum UserActionTypes {
  UPDATE = 'user/update'
}

export interface User {
  id: string
  email: string
}

export interface UserAction {
  type: UserActionTypes
  payload: User
}

const initialState = {
  id: '',
  email: '',
}

export default function userReducer(state: User = clone(initialState), action: UserAction) {
  if(action.type === UserActionTypes.UPDATE) {
    return clone(action.payload)
  }

  return state
}
