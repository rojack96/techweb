
export type Sighting = {
  id: string
  sightingId: number
  title: string
  description: string
  breed: string
  position: [number, number]
  spottedAt: number | null
}

export type Breed = {
  id: number
  name: string
}

export type CreateSightingRequest = {
  animalId: number
  breedId: number
  title: string
  description: string
  position: [number, number]
  spottedAt: number | null
}

export type CreateSightingResponse = {
  id: number
  animalId: number
  breedId: number
  title: string
  description: string
  position: [number, number]
  spottedAt: number | null
}