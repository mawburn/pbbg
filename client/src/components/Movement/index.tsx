import React, { FC, useEffect } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import isEqual from 'lodash/isEqual'
import cn from 'classnames'

import { AppState } from '../../reducers'
import { SectorsActionTypes } from '../../reducers/sectors'
import useFetch from '../../hooks/useFetch'

import arrow from './arrow.svg'
import deploy from './deploy.svg'

import s from './styles.module.scss'

const Movement: FC = () => {
  const curSector = useSelector(
    (state: AppState) => state.sectors.currentSector
  )
  const system = useSelector((state: AppState) => state.systems, isEqual)

  const dispatch = useDispatch()
  const [moveState, runFetch] = useFetch('POST', '/move')

  useEffect(() => {
    if (!moveState.isFetching && moveState.errors.length === 0 && dispatch) {
      dispatch({
        type: SectorsActionTypes.UPDATE,
        playerUpdate: moveState.response,
      })
    }
  }, [moveState, dispatch])

  const curSystem =
    system && curSector && curSector.systemId
      ? system[curSector.systemId]
      : [[]]

  return (
    <div className={s.cont}>
      <button
        disabled={!curSector || curSector.y - 1 < 0}
        className={cn(s.dir, s.upBtn)}
        onClick={() => runFetch({ direction: 'up' })}
      >
        <img src={arrow} className={s.up} alt="up" />
      </button>

      <button
        disabled={!curSector || curSector.x - 1 < 0}
        className={cn(s.dir, s.leftBtn)}
        onClick={() => runFetch({ direction: 'left' })}
      >
        <img src={arrow} className={s.left} alt="left" />
      </button>

      <button disabled className={cn(s.dir, s.deployBtn)}>
        <img src={deploy} className={s.deploy} alt="deploy" />
      </button>

      <button
        disabled={
          !curSector || curSector.x + 1 > curSystem[curSector.y].length - 1
        }
        className={cn(s.dir, s.rightBtn)}
        onClick={() => runFetch({ direction: 'right' })}
      >
        <img src={arrow} className={s.right} alt="right" />
      </button>

      <button
        disabled={!curSector || curSector.y + 1 > curSystem.length - 1}
        className={cn(s.dir, s.downBtn)}
        onClick={() => runFetch({ direction: 'down' })}
      >
        <img src={arrow} className={s.down} alt="down" />
      </button>
    </div>
  )
}

export default Movement
