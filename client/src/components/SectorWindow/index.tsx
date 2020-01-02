import React, { FC, useState, useEffect, CSSProperties } from 'react'
import cn from 'classnames'
import { useSelector } from 'react-redux'

import { AppState } from '../../reducers'
import ktn from '../../utils/keyToNumber'

import CelestialComp from './Celestial'
import SectorObjectComp from './SectorObject'

import s from './styles.module.scss'

export type SectorWindowType = 'space'

const SectorWindow: FC = () => {
  const currentSector = useSelector(
    (state: AppState) => state.sectors.currentSector
  )

  const getBg = (id: string = 'pbbg') => ({
    backgroundPosition: `${ktn(id, -1680, 1680)}px ${ktn(id, -480, 480)}px`,
    backgroundSize: `${ktn(id, 225, 800)}px`,
  })

  const [bg, setBg] = useState<CSSProperties>(getBg(''))

  useEffect(() => {
    if (currentSector) {
      setBg(getBg(currentSector.id))
    }
  }, [currentSector])

  return (
    <div style={bg} className={cn(s.cont, s.space)}>
      {currentSector && (
        <>
          {currentSector.objects.map((o: SectorObject, i: number) =>
            !o ? null : (
              <SectorObjectComp
                key={currentSector.objects[i].id}
                data={currentSector.objects[i]}
                className={s[`site${i}`]}
              />
            )
          )}
          <CelestialComp
            data={currentSector.celestial}
            className={s.celestial}
          />
        </>
      )}
    </div>
  )
}

export default SectorWindow
