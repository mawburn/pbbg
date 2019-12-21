import React, { FC, useState, useEffect, CSSProperties } from 'react'
import cn from 'classnames'

import keyToNumber from '../../utils/keyToNumber'

import s from './styles.module.scss'

export type SectorWindowType = 'space'

interface SectorWindowProps {
  sectorId: string
  type: SectorWindowType
}

const SectorWindow: FC<SectorWindowProps> = ({ sectorId, type }) => {
  const getBg = (id: string) => ({
    backgroundPosition: `${keyToNumber(id, -1522, 1522)}px ${keyToNumber(
      id,
      -435,
      435
    )}px`,
    backgroundSize: `${keyToNumber(id, 225, 800)}px`,
  })

  const [bg, setBg] = useState<CSSProperties>(getBg(sectorId))

  useEffect(() => {
    setBg(getBg(sectorId))
  }, [sectorId])

  return <div style={bg} className={cn(s.cont, s.space)} />
}

export default SectorWindow
