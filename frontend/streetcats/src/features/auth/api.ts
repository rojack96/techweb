import { httpClient } from "../../services/httpClient"
import type { ApiResponse } from "../../types/type"
import type { User } from "../../types/auth"

export function getMe(): Promise<ApiResponse<User>> {
  return httpClient<ApiResponse<User>>("/auth/me")
}