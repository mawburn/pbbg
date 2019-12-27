import React, { FC, MouseEvent } from 'react'
import cn from 'classnames'

// import randomInt from '../../../utils/randomInt'
import useMax2BgSize from './useMax2BgSize'

import s from './styles.module.scss'

export interface SectorObject {
  data: GameMap.SectorObject | null
  className: string
}

// const getClass = (): string => {
//   const _speed = randomInt(1, 30)
//   const _dir = randomInt(0, 2) === 1 ? '-rev' : ''

//   return s[`speed-${_speed}${_dir}`]
// }

const SectorObject: FC<SectorObject> = ({ data, className }) => {
  // const clssIndex = useRef<string>(getClass())
  const size = useMax2BgSize(data ? data.max : 0)

  const _handleClick = (e: MouseEvent) => {
    e.preventDefault()
  }

  return !data ? null : (
    <button
      style={{ backgroundSize: `${size}%` }}
      className={cn(s.btn, className, s[data.type])}
      onClick={_handleClick}
    />
  )
}

export default SectorObject
