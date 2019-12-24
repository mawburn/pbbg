import React, { FC, useState, useEffect } from 'react'

import useFetch from '../../hooks/useFetch'
import SectorWindow from '../SectorWindow'
import Movement from '../Movement'

import randomString from '../../utils/randomString'

import s from './styles.module.scss'

const App: FC = () => {
  const [sectorId, setSectorId] = useState<string>(randomString(10))
  
  const [res, callFetch] = useFetch('GET', '/map')
  
  useEffect(() => {
    callFetch()
  }, [callFetch])
  
  useEffect(() => {
    console.log(res)
  }, [res])

  const _handleMove = () => {
    setSectorId(randomString(10))
  }

  return (
    <div className={s.app}>
      <SectorWindow type="space" sectorId={sectorId} />
      <div />
      <Movement handleMove={_handleMove} />
    </div>
  )
}

export default App
