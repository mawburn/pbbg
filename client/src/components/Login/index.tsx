import { useEffect } from 'react'
import { useHistory, useLocation } from 'react-router-dom'

import useFetch from '../../hooks/useFetch'
import { useDispatch } from 'react-redux'

import { User, UserActionTypes } from '../../reducers/user'

const Login = () => {
  const dispatch = useDispatch()
  const location = useLocation()
  const history = useHistory()
  const [fetchState, runFetch] = useFetch<User>('POST', '/login')

  useEffect(() => {
    const { isFetching, response, errors, statusCode } = fetchState

    if (!isFetching && !response && errors.length === 0) {
      const params = new URLSearchParams(location.search)

      runFetch({ code: params.get('code') })
    } else if (statusCode === 200) {
      dispatch({ type: UserActionTypes.UPDATE, payload: response })
      history.push('/play')
    }
  }, [fetchState, location, runFetch, history, dispatch])

  return null
}

export default Login
