import React, { FC } from 'react'
import { BrowserRouter, Switch, Route, Redirect } from 'react-router-dom'

import config from './config'
import Login from './components/Login'
import App from './components/App'

const Router: FC = () => (
  <BrowserRouter>
    <Switch>
      <Route exact path="/play" component={App} />
      <Route exact path="/auth" component={() => {
          if (localStorage.getItem('token')) {
            return <Redirect to="/play" />
          }
          
          window.location.replace(config.login)
          
          return null
        }} />
      <Route exact path="/login" component={Login} />
    </Switch>
  </BrowserRouter>
)

export default Router
