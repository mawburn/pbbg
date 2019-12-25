import React, { FC, useState, useEffect, CSSProperties } from 'react'
import cn from 'classnames'
import { useSelector } from 'react-redux'
import isEqual from 'lodash/isEqual'

import { AppState } from '../../reducers'
import ktn from '../../utils/keyToNumber'

import Celestial from './Celestial'
import SectorObject from './SectorObject'

import s from './styles.module.scss'

export type SectorWindowType = 'space'

interface SectorWindowProps {
  sectorId: string
  type: SectorWindowType
}

const SectorWindow: FC<SectorWindowProps> = ({ sectorId, type }) => {
  const curSector = useSelector((state: AppState) => state.sector, isEqual)
  const sectors = useSelector(
    (state: AppState) => state.map.systems[0].sectors,
    isEqual
  )

  const [sector, setSector] = useState<GameMap.Sector | null>(null)

  const getBg = (id: string) => ({
    backgroundPosition: `${ktn(id, -1680, 1680)}px ${ktn(id, -480, 480)}px`,
    backgroundSize: `${ktn(id, 225, 800)}px`,
  })

  const [bg, setBg] = useState<CSSProperties>(getBg(''))

  useEffect(() => {
    const _sector = sectors[curSector.y][curSector.x]

    if (_sector) {
      setBg(getBg(_sector.id))
      setSector(_sector)
    }
  }, [curSector, sectors])

  return (
    <div style={bg} className={cn(s.cont, s.space)}>
      {sector && (
        <>
          {sector.objects.map((o, i) =>
            !o ? null : (
              <SectorObject
                key={sector.objects[i].id}
                data={sector.objects[i]}
                className={s[`site${i}`]}
              />
            )
          )}
          <Celestial data={sector.celestial} className={s.celestial} />
        </>
      )}
    </div>
  )
}

export default SectorWindow
