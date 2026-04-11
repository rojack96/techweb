import { httpClient } from "../../services/httpClient"
import type { Sighting } from "../../types/sightings"
import type { ApiResponse } from "../../types/type"

const CAT = "1"

export function getSightings(): Promise<ApiResponse<Sighting[]>> {
  return httpClient<ApiResponse<Sighting[]>>(`/sightings/${CAT}/all`)
}