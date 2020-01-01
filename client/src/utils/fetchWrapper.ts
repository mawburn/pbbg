import FetchError from './FetchError'

import randomString from './randomString'

export type FetchMethods = 'GET' | 'POST' | 'PUT' | 'DEL'

const userToken = randomString(15)

/**
 * Performs a JSON stringify on a body for the fetch request
 * Private Function
 */
const bodyToString = (body?: string | object): string | undefined =>
  body === undefined
    ? undefined
    : typeof body === 'string'
    ? body
    : JSON.stringify(body)

/**
 * Generic Fetch Wrapper
 */
const fetchWrapper = (
  method: FetchMethods,
  endpoint: string,
  body?: object | string,
  opts: any = {}
) => {
  switch (method) {
    case 'GET':
      return fetchGet(endpoint, body, '', opts)
    case 'POST':
      return fetchPost(endpoint, body, '', opts)
    case 'PUT':
      return fetchPut(endpoint, body, '', opts)
    case 'DEL':
      return fetchDelete(endpoint, body, '', opts)
  }
}

/**
 * GET Request
 */
const fetchGet = (
  endpoint: string,
  body?: object | string,
  contentType?: string,
  opts: any = {},
  ip?: string
) => {
  opts.method = 'GET'

  const tmpBody = typeof body === 'string' ? JSON.parse(body) : body

  const paramArr = !tmpBody
    ? []
    : Object.keys(tmpBody).map(key => `${key}=${tmpBody[key]}`)

  const bodyAsUrlParams = paramArr.length > 0 ? `?${paramArr.join('&')}` : ''

  return fetchWithErrors(endpoint + bodyAsUrlParams, opts, contentType, ip)
}

/**
 * POST Request
 */
const fetchPost = (
  endpoint: string,
  body?: string | object,
  contentType?: string,
  opts: any = {},
  ip?: string
) => {
  opts.method = 'POST'
  opts.body = bodyToString(body)

  return fetchWithErrors(endpoint, opts, contentType, ip)
}

/**
 * DELETE Request
 */
const fetchDelete = (
  endpoint: string,
  body?: string | object,
  contentType?: string,
  opts: any = {},
  ip?: string
) => {
  opts.method = 'DELETE'
  opts.body = bodyToString(body)

  return fetchWithErrors(endpoint, opts, contentType, ip)
}

/**
 * PUT Request
 */
const fetchPut = (
  endpoint: string,
  body?: string | object,
  contentType?: string,
  opts: any = {},
  ip?: string
) => {
  opts.method = 'PUT'
  opts.body = bodyToString(body)

  return fetchWithErrors(endpoint, opts, contentType, ip)
}

/**
 * Perform an fetch request, throwing on errors
 */
const fetchWithErrors = (
  endpoint: string,
  opts?: any,
  contentType?: string,
  ip?: string
) => {
  const headers: HeadersInit = {}

  headers['Content-Type'] = contentType || 'application/json'

  if (userToken) {
    headers['Authorization'] = `Bearer ${userToken}`
  }

  if (ip) {
    headers['X-Forwarded-For'] = ip
  }

  const fetchOpts: any = {
    ...opts,
    credentials: 'include',
    headers: new Headers({ ...headers }),
  }

  return fetch(endpoint, fetchOpts)
    .then(res => {
      const contentType = res.headers.get('Content-Type')

      const cType =
        contentType && contentType.includes('json')
          ? 'json'
          : contentType && contentType.includes('text')
          ? 'text'
          : 'unk'

      const statusCode: number = res.status

      if (!res.ok) {
        if (statusCode === 401) {
          throw new FetchError('Unauthorized', statusCode)
        } else if (cType !== 'unk') {
          if (cType === 'json') {
            return res.json().then(data => {
              const errText = data.error || JSON.stringify(data)

              throw new FetchError(errText, statusCode)
            })
          } else if (cType === 'text') {
            return res.text().then(data => {
              throw new FetchError(data, statusCode)
            })
          }
        }

        // If we haven't already thrown, throw status code
        throw new FetchError('Unknown Error', statusCode)
      }

      if (cType === 'json') {
        return res.json()
      } else if (cType === 'text') {
        return res.text()
      } else {
        return res
      }
    })
    .catch(err => {
      if (err.message.toLowerCase() === 'failed to fetch') {
        throw new FetchError('Failed')
      } else if (err.message.includes('Logged Out')) {
        console.warn('Logged out')
      } else {
        throw err
      }
    })
}

export {
  fetchWithErrors,
  fetchWrapper,
  fetchGet,
  fetchPost,
  fetchPut,
  fetchDelete,
}
