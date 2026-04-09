
import { httpClient } from "../../services/httpClient"
import type { Sighting } from "./types"

const CAT = "1"

export function getSightings(): Promise<Sighting[]> {
  return httpClient<Sighting[]>(`/sightings/${CAT}/all`)
}