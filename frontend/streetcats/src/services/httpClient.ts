const BASE_URL = import.meta.env.VITE_API_URL

export async function httpClient<T>(
  endpoint: string,
  options?: RequestInit
): Promise<T> {
  const res = await fetch(`${BASE_URL}${endpoint}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      ...options?.headers,
    },
  })

  if (!res.ok) {
    throw new Error("API error")
  }

  return res.json()
}