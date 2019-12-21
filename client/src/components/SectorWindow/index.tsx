import React, { FC, useRef, CSSProperties } from 'react'
import cn from 'classnames'

import keyToNumber from '../../utils/keyToNumber'

import s from './styles.module.scss'

export type SectorWindowType = 'space'

interface SectorWindowProps {
  sectorId: string
  type: SectorWindowType
}

const SectorWindow: FC<SectorWindowProps> = ({ sectorId, type }) => {
  const horiz = useRef<number>(keyToNumber(sectorId, -1522, 1522))
  const vert = useRef<number>(keyToNumber(sectorId, -435, 435))
  const zoom = useRef<number>(keyToNumber(sectorId, 200, 800))

  const background: CSSProperties = {
    backgroundPosition: `${horiz.current}px ${vert.current}px`,
    backgroundSize: `${zoom.current}px`,
  }

  return <div style={background} className={cn(s.cont, s.space)} />
}

export default SectorWindow
