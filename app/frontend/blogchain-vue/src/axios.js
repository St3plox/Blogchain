import axios from 'axios';
import { getToken } from '@/utils/auth';

const instance = axios.create({
  baseURL: process.env.VUE_APP_BACKEND_URL || 'http://localhost:3000/v1',
  headers: {
    'Content-Type': 'application/json',
  },
});
instance.interceptors.request.use(config => {
  const token = getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default instance;
