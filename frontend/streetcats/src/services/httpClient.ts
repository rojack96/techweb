const BASE_URL = import.meta.env.VITE_API_URL

export async function httpClient<T>(
  endpoint: string,
  options?: RequestInit
): Promise<T> {
  const token = localStorage.getItem("token")
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...((options?.headers ?? {}) as Record<string, string>),
  }

  if (token) {
    headers.Authorization = `Bearer ${token}`
  }

  const res = await fetch(`${BASE_URL}${endpoint}`, {
    credentials: "include",
    ...options,
    headers,
  })

  if (!res.ok) {
    throw new Error("API error")
  }

  return res.json()
}