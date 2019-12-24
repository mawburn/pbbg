/**
 * Custom Error for handling problems with Fetch API Calls
 */
export default class FetchError extends Error {
  statusCode?: number

  constructor(message?: string, statusCode?: number) {
    super(message)
    this.name = 'Fetch Error'
    this.statusCode = statusCode
  }
}
