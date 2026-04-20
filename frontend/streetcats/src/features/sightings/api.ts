import { httpClient } from "../../services/httpClient"
import type { Breed, CreateSightingRequest, CreateSightingResponse, Sighting } from "../../types/sightings"
import type { ApiResponse } from "../../types/type"

const CAT = "1"

export function getSightings(): Promise<ApiResponse<Sighting[]>> {
  return httpClient<ApiResponse<Sighting[]>>(`/sightings/${CAT}/all`)
}

export function getBreedsLookup(): Promise<ApiResponse<Breed[]>> {
  return httpClient<ApiResponse<Breed[]>>(`/sightings/${CAT}/breeds-lookup`)
}

export function createSighting(data: CreateSightingRequest): Promise<ApiResponse<CreateSightingResponse>> {
  return httpClient<ApiResponse<CreateSightingResponse>>("/sightings/create", {
    method: "POST",
    body: JSON.stringify(data),
  })
}