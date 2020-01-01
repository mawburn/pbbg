import { useEffect, useRef, useState } from 'react'
import { useHistory } from 'react-router-dom'

import config from '../config'
import { FetchMethods, fetchWrapper } from '../utils/fetchWrapper'

export interface FetchState<T> {
  isFetching: boolean
  errors: string[]
  response: T | null
}

const initialState = {
  isFetching: false,
  errors: [],
  response: null,
}

/**
 * @param method http method
 * @param url api url to fetch, will use `config.api` settings
 *
 * @returns {[FetchState<T>, (requestPayload: Object) => void]} api response
 */
const useFetch = <T>(
  method: FetchMethods,
  url: string
): [FetchState<T>, (requestPayload?: Object) => void] => {
  const history = useHistory()
  const [state, setState] = useState<FetchState<T>>({ ...initialState })

  const controller = useRef(new AbortController())

  const callFetch = async (requestPayload?: Object) => {
    if (!state.isFetching) {
      setState({ ...state, isFetching: true })

      try {
        const opts = {
          signal: controller.current.signal,
        }

        const response = await fetchWrapper(
          method,
          `${config.api}${url}`,
          requestPayload,
          opts
        )

        setState({ isFetching: false, errors: [], response })
      } catch (err) {
        if (err.name === 'Fetch Error') {
          history.replace('/auth')
        }

        console.error(err)
        const errors = Array.isArray(err) ? [...err] : [err]
        setState({ isFetching: false, errors, response: null })
      }
    }
  }

  useEffect((): any => (): void => controller.current.abort(), [])

  return [state, callFetch]
}

export default useFetch
