export default function keyToNumber(key: string, min: number, max: number) {
  const keyNum = key.split('').reduce((acc, c) => acc + c.charCodeAt(0), 0)

  const base = Math.sin(keyNum) * 1000 * key.length * min

  return Math.round((base - Math.floor(base)) * (max - min) + min)
}
