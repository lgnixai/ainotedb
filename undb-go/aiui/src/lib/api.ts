import { Project } from '../types';
import { generateId } from './utils';
import axios from 'axios';
import { Project } from '../types';
import { generateId } from './utils';

//const API_BASE = 'http://localhost:5555/api'; // 可根据 golang 后端实际路径调整
const API_BASE = '/api';

const api = axios.create({
  baseURL: API_BASE,
  timeout: 10000,
  // 如果后端需要携带 cookie，解开下面一行注释
  // withCredentials: true,
});

// 添加请求拦截器，自动携带 token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers = config.headers || {};
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// 注册
export async function register(email: string, password: string) {
  try {
    const res = await api.post('/users/register', { email, password });
    // 保存 token 和 user
    if (res.data.token) {
      localStorage.setItem('token', res.data.token);
    }
    if (res.data.user) {
      localStorage.setItem('user', JSON.stringify(res.data.user));
    }
    return res.data;
  } catch (error: any) {
    throw new Error(error.response?.data?.message || '注册失败');
  }
}

// 登录
export async function login(email: string, password: string) {
  try {
    const res = await api.post('/users/login', { email, password });
    // 保存 token 和 user
    if (res.data.token) {
      localStorage.setItem('token', res.data.token);
    }
    if (res.data.user) {
      localStorage.setItem('user', JSON.stringify(res.data.user));
    }
    return res.data;
  } catch (error: any) {
    throw new Error(error.response?.data?.message || '登录失败');
  }
}

// 创建团队空间（space）
export async function createSpace({ name, description = '', visibility = 'public' }: { name: string; description?: string; visibility?: string }) {
  try {
    const res = await api.post('/spaces', { name, description, visibility });
    return res.data;
  } catch (error: any) {
    throw new Error(error.response?.data?.message || '空间创建失败');
  }
}


export async function getSpaces() {
  const res = await api.get('/spaces');
  return res.data;
}

// 获取某个 space 下所有表
export async function getTablesBySpaceId(spaceId: string) {
  const res = await api.get(`/tables/space/${spaceId}`);
  return res.data;
}

// 创建新表（table）
export async function createTable(spaceId: string | number, name: string) {
  const res = await api.post('/tables', { space_id:  (spaceId), name });
  return res.data;
}
