import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';

const instance: AxiosInstance = axios.create({
  baseURL: process.env.API_URL,
});

instance.interceptors.request.use((config: AxiosRequestConfig) => {
  // TODO: set authorization token

  return config;
});

export default instance;
