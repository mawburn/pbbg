import React, { FC, useState, useEffect } from 'react'
import { useDispatch } from 'react-redux'

import useFetch from '../../hooks/useFetch'
import SectorWindow from '../SectorWindow'
import Movement from '../Movement'
import { SectorsActionTypes, Sectors } from '../../reducers/sectors'
import { SystemsActionTypes, Systems } from '../../reducers/systems'

import randomString from '../../utils/randomString'

import s from './styles.module.scss'

interface MapResponse {
  systems: Systems
  sectors: Sectors
}

const App: FC = () => {
  const [sectorId, setSectorId] = useState<string>(randomString(10))
  const [fetchState, callFetch] = useFetch<MapResponse>('GET', '/map')
  const dispatch = useDispatch()

  useEffect(() => {
    if (
      !fetchState.response &&
      !fetchState.isFetching &&
      fetchState.errors.length === 0
    ) {
      callFetch()
    }
  }, [callFetch, fetchState])

  useEffect(() => {
    if (!fetchState.isFetching && fetchState.response && dispatch) {
      const { systems, sectors } = fetchState.response

      dispatch({
        type: SystemsActionTypes.LOAD,
        payload: systems,
      })

      dispatch({
        type: SectorsActionTypes.LOAD,
        payload: sectors,
      })
    }
  }, [fetchState, dispatch])

  const _handleMove = () => {
    setSectorId(randomString(10))
  }

  return (
    <div className={s.app}>
      {!fetchState.response ? (
        <div>Loading...</div>
      ) : (
        <>
          <SectorWindow type="space" sectorId={sectorId} />
          <div />
          <Movement handleMove={_handleMove} />
        </>
      )}
    </div>
  )
}

export default App
