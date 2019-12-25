export default function randomInt(min: number, max: number): number {
  const _min = Math.ceil(min)
  const _max = Math.floor(max)
  return Math.floor(Math.random() * (_max - _min)) + _min
}
