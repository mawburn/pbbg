import { useEffect, useState } from 'react'

import keyToNumber from '../utils/keyToNumber'

const useKey2Num = (key: string, min: number, max: number): number => {
  const [num, setNum] = useState<number>(keyToNumber(key, min, max))
  
  useEffect(() => {
    setNum(keyToNumber(key, min, max))
  }, [key, min, max])
  
  return num
}

export default useKey2Num