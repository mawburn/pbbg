import React, { FC, useRef, CSSProperties } from 'react'
import cn from 'classnames'

import randomNumber from '../../utils/randomNumber'
import s from './styles.module.scss'

export type SectorWindowType = 'space'

interface SectorWindowProps {
  type: SectorWindowType
}

const SectorWindow: FC<SectorWindowProps> = ({ type }) => {
  // TODO I want to base this on a seed so I get the same numbers every time for a given sector key
  const horiz = useRef<number>(randomNumber(-1522, 1522))
  const vert = useRef<number>(randomNumber(-435, 435))
  const zoom = useRef<number>(randomNumber(200, 800))
  
  const background: CSSProperties = {
    backgroundPosition: `${horiz.current}px ${vert.current}px`,
    backgroundSize: `${zoom.current}px`,
  }
  
  console.log(background)

  return <div style={background} className={cn(s.cont, s.space)}  />
}

export default SectorWindow
