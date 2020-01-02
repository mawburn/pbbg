import React, { FC, useState, useEffect, CSSProperties } from 'react'
import cn from 'classnames'
import { useSelector } from 'react-redux'
import isEqual from 'lodash/isEqual'

import { AppState } from '../../reducers'
import ktn from '../../utils/keyToNumber'

import CelestialComp from './Celestial'
import SectorObjectComp from './SectorObject'

import s from './styles.module.scss'

export type SectorWindowType = 'space'

interface SectorWindowProps {
  sectorId: string
  type: SectorWindowType
}

const SectorWindow: FC<SectorWindowProps> = ({ sectorId, type }) => {
  const playerSector = useSelector((state: AppState) => state.playerSector)
  const sectors = useSelector(
    (state: AppState) => state.sectors,
    isEqual
  )

  const [sector, setSector] = useState<Sector | null>(null)

  const getBg = (id: string) => ({
    backgroundPosition: `${ktn(id, -1680, 1680)}px ${ktn(id, -480, 480)}px`,
    backgroundSize: `${ktn(id, 225, 800)}px`,
  })

  const [bg, setBg] = useState<CSSProperties>(getBg(''))

  useEffect(() => {
    const _sector = sectors[playerSector.id]

    if (_sector) {
      setBg(getBg(playerSector.id))
      setSector(_sector)
    }
  }, [playerSector, sectors])

  return (
    <div style={bg} className={cn(s.cont, s.space)}>
      {sector && (
        <>
          {sector.objects.map((o: SectorObject, i: number) =>
            !o ? null : (
              <SectorObjectComp
                key={sector.objects[i].id}
                data={sector.objects[i]}
                className={s[`site${i}`]}
              />
            )
          )}
          <CelestialComp data={sector.celestial} className={s.celestial} />
        </>
      )}
    </div>
  )
}

export default SectorWindow
