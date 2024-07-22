// src/axios.js
import axios from 'axios';
import { getToken } from '@/utils/auth';

const instance = axios.create({
  baseURL: 'http://localhost:3000', // Change to your backend URL
});

instance.interceptors.request.use(config => {
  const token = getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default instance;
