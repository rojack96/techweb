// types.ts
export interface ApiResponse<T> {
  message: string;
  response: T;
  status: number;
}