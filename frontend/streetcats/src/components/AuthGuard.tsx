import { useAuth } from '../hooks/auth'

export const AuthGuard = ({ children }: { children: React.ReactNode }) => {
    const { isLogged, loading } = useAuth()
    console.log("AuthGuard", { isLogged, loading })
    if (loading) return null

    if (!isLogged) return null

    return <>{children}</>
}