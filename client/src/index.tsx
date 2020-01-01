import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import { createBrowserHistory } from 'history'

import './index.scss'
import configureStore from './store'
import Router from './Router'

export const history = createBrowserHistory()

const store = configureStore()

ReactDOM.render(
  <Provider store={store}>
    <Router />
  </Provider>,
  document.getElementById('root')
)
