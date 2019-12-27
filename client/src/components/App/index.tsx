import React, { FC, useState, useEffect } from 'react'
import { useDispatch } from 'react-redux'

import useFetch from '../../hooks/useFetch'
import SectorWindow from '../SectorWindow'
import Movement from '../Movement'
import { MapActions } from '../../reducers/map'

import randomString from '../../utils/randomString'

import s from './styles.module.scss'

const App: FC = () => {
  const [sectorId, setSectorId] = useState<string>(randomString(10))
  const [fetchState, callFetch] = useFetch('GET', '/map')
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
      dispatch({
        type: MapActions.LOAD,
        payload: fetchState.response,
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
