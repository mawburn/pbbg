import React, { FC } from 'react'
import cn from 'classnames'

import planet from './imgs/planet.png'
import station from './imgs/station.png'
import s from './styles.module.scss'

export interface Celestial {
  data: GameMap.Celestial | null
  className: string
}

const Celestial: FC<Celestial> = ({ data, className }) =>
  !data ? null : (
    <button className={cn(s.cont, className)}>
      <img src={data.type === 'planet' ? planet : station} alt={data.type} />
    </button>
  )

export default Celestial
