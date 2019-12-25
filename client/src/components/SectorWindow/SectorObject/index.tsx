import React, { FC, MouseEvent, useRef } from 'react'
import cn from 'classnames'

import randomInt from '../../../utils/randomInt'

import ore from './imgs/ore.png'
import s from './styles.module.scss'

export interface SectorObject {
  data: GameMap.SectorObject | null
  className: string
}

const getClass = (): string => {
  const _speed = randomInt(1, 30)
  const _dir = randomInt(0, 2) === 1 ? '-rev' : ''

  return s[`speed-${_speed}${_dir}`]
}

const SectorObject: FC<SectorObject> = ({ data, className }) => {
  const clssIndex = useRef<string>(getClass())

  const _handleClick = (e: MouseEvent) => {
    e.preventDefault()
  }

  return !data ? null : (
    <button className={cn(s.btn, className)} onClick={_handleClick}>
      <img src={ore} className={clssIndex.current} alt={data.type} />
    </button>
  )
}

export default SectorObject
