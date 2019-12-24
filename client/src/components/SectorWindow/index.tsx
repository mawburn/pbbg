import React, { FC, useState, useEffect, CSSProperties } from 'react'
import cn from 'classnames'
import { useSelector } from 'react-redux'
import isEqual from 'lodash/isEqual'

import { AppState } from '../../reducers'
import ktn from '../../utils/keyToNumber'

import s from './styles.module.scss'

export type SectorWindowType = 'space'

interface SectorWindowProps {
  sectorId: string
  type: SectorWindowType
}

const SectorWindow: FC<SectorWindowProps> = ({ sectorId, type }) => {
  const curSector = useSelector((state: AppState) => state.sector, isEqual)
  const sectors = useSelector((state: AppState) => state.map.systems[0].sectors, isEqual)
  
  const getBg = (id: string) => ({
    backgroundPosition: `${ktn(id, -1680, 1680)}px ${ktn(id, -480, 480)}px`,
    backgroundSize: `${ktn(id, 225, 800)}px`,
  })

  const [bg, setBg] = useState<CSSProperties>(getBg(''))

  useEffect(() => {
    const sector = sectors[curSector.y][curSector.x]
    
    if (sector) {
      setBg(getBg(sector.id))
    }
  }, [curSector])

  return <div style={bg} className={cn(s.cont, s.space)} />
}

export default SectorWindow
