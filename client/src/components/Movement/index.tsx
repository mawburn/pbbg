import React, { FC } from 'react'

import arrow from './arrow.svg'
import deploy from './deploy.svg'

import s from './styles.module.scss'

const Movement: FC = () => (
  <div className={s.cont}>
    <div className={s.placeholder} />
    <button className={s.dir}>
      <img src={arrow} className={s.up} alt="up" />
    </button>
    <div className={s.placeholder} />

    <button className={s.dir}>
      <img src={arrow} className={s.left} alt="left" />
    </button>
    <button className={s.dir}>
      <img src={deploy} className={s.deploy} alt="deploy" />
    </button>
    <button className={s.dir}>
      <img src={arrow} className={s.right} alt="right" />
    </button>

    <div className={s.placeholder} />
    <button className={s.dir}>
      <img src={arrow} className={s.down} alt="down" />
    </button>
    <div className={s.placeholder} />
  </div>
)

export default Movement
