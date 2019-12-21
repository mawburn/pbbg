import React, { FC } from 'react'

import SectorWindow from '../SectorWindow'
import Movement from '../Movement'

import s from './styles.module.scss'

const App: FC = () => (
  <div className={s.app}>
    <SectorWindow type="space" />
    <Movement />
  </div>
)

export default App
