import { useQuery } from "@tanstack/react-query"
import { getSightings } from "./api"

export function useSightings() {
  return useQuery({
    queryKey: ["sightings"],
    queryFn: getSightings,
  })
}