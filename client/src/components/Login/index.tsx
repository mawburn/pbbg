import { useEffect } from 'react'
import { useHistory, useLocation } from 'react-router-dom'

import useFetch from '../../hooks/useFetch'

const Login = () => {
  const location = useLocation()
  const history = useHistory()
  const [fetchState, runFetch] = useFetch<{ ok: boolean }>('POST', '/login')

  useEffect(() => {
    const { isFetching, response, errors } = fetchState

    if (!isFetching && !response && errors.length === 0) {
      const params = new URLSearchParams(location.search)

      runFetch({ code: params.get('code') })
    } else if (response && response.ok) {
      history.push('/play')
    }
  }, [fetchState, location, runFetch, history])

  return null
}

export default Login
