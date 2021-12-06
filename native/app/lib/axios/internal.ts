import axios, { AxiosInstance } from 'axios';

const internalInstance: AxiosInstance = axios.create({
  baseURL: process.env.API_URL,
});

export default internalInstance;
