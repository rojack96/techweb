import { useMutation, useQueryClient, useQuery } from "@tanstack/react-query"
import { createSighting, getBreedsLookup, getSightings } from "./api"

export function useSightings() {
  return useQuery({
    queryKey: ["sightings"],
    queryFn: getSightings,
  })
}


export function useBreedsLookup() {
  return useQuery({
    queryKey: ["breeds", "lookup"],
    queryFn: getBreedsLookup,
    select: (data) => data.response
  })
}


export function useCreateSighting() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: createSighting,

    onSuccess: () => {
      // invalida la lista per rifetch
      queryClient.invalidateQueries({ queryKey: ["sightings"] })
    }
  })
}