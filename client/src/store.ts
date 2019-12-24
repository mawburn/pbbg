import { compose, createStore } from 'redux'

import rootReducer, { AppState } from './reducers'

export default function configureStore(preloadedState?: AppState) {
  const composeEnhancer: typeof compose = (window as any)
    .__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
    ? (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({
        trace: true,
        traceLimit: 25,
      })
    : compose

  const store = createStore(rootReducer(), composeEnhancer())

  // Hot reloading
  if (module.hot) {
    // Enable Webpack hot module replacement for reducers
    module.hot.accept('./reducers', () => {
      store.replaceReducer(rootReducer())
    })
  }

  return store
}
