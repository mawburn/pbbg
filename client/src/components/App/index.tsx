import React, { FC, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'

import useFetch from '../../hooks/useFetch'
import SectorWindow from '../SectorWindow'
import Movement from '../Movement'
import { AppState } from '../../reducers'
import { SectorsActionTypes, Sectors } from '../../reducers/sectors'
import { SystemsActionTypes, Systems } from '../../reducers/systems'

import s from './styles.module.scss'

interface MapResponse {
  systems: Systems
  sectors: Sectors
}

const App: FC = () => {
  const currentSector = useSelector(
    (state: AppState) => state.sectors.currentSector
  )
  const [fetchState, callFetch] = useFetch<MapResponse>('GET', '/map')
  const [fetchSector, callFetchSector] = useFetch<Sector>(
    'GET',
    '/currentSector'
  )
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

  useEffect(() => {
    const { isFetching, errors } = fetchSector

    if (
      currentSector === null &&
      !isFetching &&
      errors.length === 0 &&
      !fetchState.isFetching &&
      fetchState.errors.length === 0 &&
      fetchState.response
    ) {
      callFetchSector()
    }
  }, [currentSector, fetchSector, callFetchSector, fetchState])

  useEffect(() => {
    const { isFetching, errors, response } = fetchSector

    if (!isFetching && errors.length === 0 && response) {
      dispatch({ type: SectorsActionTypes.UPDATE, playerUpdate: response })
    }
  }, [fetchSector, dispatch])

  return (
    <div className={s.app}>
      {!fetchState.response ? (
        <div>Loading...</div>
      ) : (
        <>
          <SectorWindow />
          <div />
          <Movement />
        </>
      )}
    </div>
  )
}

export default App
