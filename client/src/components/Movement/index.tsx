import React, { FC } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import isEqual from 'lodash/isEqual'
import cn from 'classnames'

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
        if (y - 1 >= 0) {
          dispatch({
            type: SectorActions.MOVE,
            payload: { ...curSector, y: y - 1 },
          })
        }
        return
      case 'down':
        if (y + 1 < sectors.length) {
          dispatch({
            type: SectorActions.MOVE,
            payload: { ...curSector, y: y + 1 },
          })
        }
        return
      case 'right':
        if (x + 1 < sectors[y].length) {
          dispatch({
            type: SectorActions.MOVE,
            payload: { ...curSector, x: x + 1 },
          })
        }
        return
      case 'left':
        if (x - 1 >= 0) {
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
      <button
        disabled={curSector.y - 1 < 0}
        className={cn(s.dir, s.upBtn)}
        onClick={() => handleMove('up')}
      >
        <img src={arrow} className={s.up} alt="up" />
      </button>

      <button
        disabled={curSector.x - 1 < 0}
        className={cn(s.dir, s.leftBtn)}
        onClick={() => handleMove('left')}
      >
        <img src={arrow} className={s.left} alt="left" />
      </button>

      <button disabled className={cn(s.dir, s.deployBtn)}>
        <img src={deploy} className={s.deploy} alt="deploy" />
      </button>

      <button
        disabled={curSector.x + 1 > sectors[curSector.y].length - 1}
        className={cn(s.dir, s.rightBtn)}
        onClick={() => handleMove('right')}
      >
        <img src={arrow} className={s.right} alt="right" />
      </button>

      <button
        disabled={curSector.y + 1 > sectors.length - 1}
        className={cn(s.dir, s.downBtn)}
        onClick={() => handleMove('down')}
      >
        <img src={arrow} className={s.down} alt="down" />
      </button>
    </div>
  )
}

export default Movement
