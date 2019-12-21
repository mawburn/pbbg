import React, { FC } from 'react'

import SectorWindow from '../SectorWindow'
import Movement from '../Movement'

import randomString from '../../utils/randomString'

import s from './styles.module.scss'

const App: FC = () => (
  <div className={s.app}>
    <SectorWindow type="space" sectorId={randomString(10)} />
    <Movement />
  </div>
)

export default App
