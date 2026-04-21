import { useQuery } from '@tanstack/react-query'
import { getMe } from '../features/auth/api'

export const useAuth = () => {
  const query = useQuery({
    queryKey: ['me'],
    queryFn: getMe,
    retry: false
  })

  return {
    user: query.data?.response ?? null,
    isLogged: !!query.data?.response,
    loading: query.isLoading
  }
}