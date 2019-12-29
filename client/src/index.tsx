import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'

import './index.scss'
import configureStore from './store'
import Router from './Router'

const store = configureStore()

ReactDOM.render(
  <Provider store={store}>
    <Router />
  </Provider>,
  document.getElementById('root')
)
