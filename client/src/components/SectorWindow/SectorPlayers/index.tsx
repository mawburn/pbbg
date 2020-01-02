import React, { FC, MouseEvent } from 'react'
import cn from 'classnames'

import s from './styles.module.scss'

interface SectorPlayer {
  className: string
}

const SectorPlayer: FC<SectorPlayer> = ({ className }) => {
  const _handleClick = (e: MouseEvent) => {
    e.preventDefault()
  }

  return <button className={cn(s.btn, className)} onClick={_handleClick} />
}

export default SectorPlayer
