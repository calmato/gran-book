import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import * as LocalStorage from '~/lib/local-storage';
import { Auth } from '~/store/models';

const instance: AxiosInstance = axios.create({
  baseURL: process.env.API_URL,
});

instance.interceptors.request.use(async (config: AxiosRequestConfig) => {
  const auth: Auth.Model = await LocalStorage.AuthStorage.retrieve();
  if (auth) {
    const token = `Bearer ${auth.token}`;
    config.headers.Authorization = token;
  }

  return config;
});

export default instance;
