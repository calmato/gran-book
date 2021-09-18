import { AxiosRequestConfig } from 'axios';

export function getAuthHeader(token: string): AxiosRequestConfig {
  return {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
}
