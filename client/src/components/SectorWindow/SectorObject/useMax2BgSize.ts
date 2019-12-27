import useKey2Num from '../../../hooks/useKey2Num'

const useMax2BgSize = (max: number) => { 
  let _min, _max
  
  if (max === 0) {
    _min = 0
    _max = 0
  } else if (max > 1000000) {
    _min = 60
    _max = 80
  } else if (max > 100000) {
    _min = 40
    _max = 79
  } else {
    _min = 25
    _max = 39
  }
  
  return useKey2Num(`${max}`, _min, _max)
}

export default useMax2BgSize