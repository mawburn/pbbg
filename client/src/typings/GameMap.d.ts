namespace GameMap {
  export interface SectorObject {
    id: string
    type: string
    max: number
    quantity: number
  }

  export interface Celestial {
    id: string
    type: string
    name: string
  }

  export interface Sector {
    id: string
    celestial: Celestial
    objects: SectorObject[]
  }

  export interface System {
    id: string
    sectors: Sector[][]
  }

  export interface Coords {
    [key: string]: {
      system: number
      x: number
      y: number
    }
  }

  export interface Map {
    coords: Coords
    systems: System[]
  }
}
