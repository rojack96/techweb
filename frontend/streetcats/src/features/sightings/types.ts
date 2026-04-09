
export type Sighting = {
  id: string
  sightingId: number
  title: string
  description: string
  breed: string
  position: [number, number]
  spottedAt: number | null
}