import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';

const externalInstance: AxiosInstance = axios.create({});

externalInstance.interceptors.request.use(async (config: AxiosRequestConfig) => {
  return config;
});

export default externalInstance;
