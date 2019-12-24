namespace GameMap {
  export interface SectorObject {
    id: string
    type: string
    size: string
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

  export interface Map {
    systems: System[]
  }
}
