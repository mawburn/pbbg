import React, { FC } from 'react'
import cn from 'classnames'

import useKey2Num from '../../../hooks/useKey2Num'

import s from './styles.module.scss'

export interface Celestial {
  data: GameMap.Celestial | null
  className: string
}

const Celestial: FC<Celestial> = ({ data, className }) => {
  const size = useKey2Num(data ? data.id : '', 30, 75)

  return !data ? null : (
    <button
      style={{ backgroundSize: `${size}%` }}
      className={cn(s.cont, className, s[data.type])}
    />
  )
}

export default Celestial
