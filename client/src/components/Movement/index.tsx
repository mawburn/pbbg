import React, { FC } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import isEqual from 'lodash/isEqual'

import { AppState } from '../../reducers'
import { SectorActions } from '../../reducers/sector'

import arrow from './arrow.svg'
import deploy from './deploy.svg'

import s from './styles.module.scss'

export type MoveDirection = 'up' | 'down' | 'left' | 'right'

interface MovementProps {
  handleMove: (dir: MoveDirection) => void
}

const Movement: FC<MovementProps> = () => {
  const curSector = useSelector((state: AppState) => state.sector, isEqual)
  const sectors = useSelector(
    (state: AppState) => state.map.systems[0].sectors,
    isEqual
  )
  const dispatch = useDispatch()

  const handleMove = (d: MoveDirection) => {
    const { x, y } = curSector

    switch (d) {
      case 'up':
        if (sectors[y + 1][x]) {
          dispatch({
            type: SectorActions.MOVE,
            payload: { ...curSector, y: y + 1 },
          })
        }
        return
      case 'down':
        if (sectors[y - 1][x]) {
          dispatch({
            type: SectorActions.MOVE,
            payload: { ...curSector, y: y - 1 },
          })
        }
        return
      case 'right':
        if (sectors[y][x + 1]) {
          dispatch({
            type: SectorActions.MOVE,
            payload: { ...curSector, x: x + 1 },
          })
        }
        return
      case 'left':
        if (sectors[y][x - 1]) {
          dispatch({
            type: SectorActions.MOVE,
            payload: { ...curSector, x: x - 1 },
          })
        }
        return
    }
  }

  return (
    <div className={s.cont}>
      <div className={s.placeholder} />
      <button className={s.dir} onClick={() => handleMove('up')}>
        <img src={arrow} className={s.up} alt="up" />
      </button>
      <div className={s.placeholder} />

      <button className={s.dir} onClick={() => handleMove('left')}>
        <img src={arrow} className={s.left} alt="left" />
      </button>
      <button className={s.dir}>
        <img src={deploy} className={s.deploy} alt="deploy" />
      </button>
      <button className={s.dir} onClick={() => handleMove('right')}>
        <img src={arrow} className={s.right} alt="right" />
      </button>

      <div className={s.placeholder} />
      <button className={s.dir} onClick={() => handleMove('down')}>
        <img src={arrow} className={s.down} alt="down" />
      </button>
      <div className={s.placeholder} />
    </div>
  )
}

export default Movement
