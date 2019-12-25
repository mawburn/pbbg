import React, { FC } from 'react'
import cn from 'classnames'

import s from './styles.module.scss'

export interface Celestial {
  data: GameMap.Celestial | null
  className: string
}

const Celestial: FC<Celestial> = ({ data, className }) =>
  !data ? null : <button className={cn(s.cont, className, s[data.type])} />

export default Celestial
