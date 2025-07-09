import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';

// 创建axios实例
const request: AxiosInstance = axios.create({
  baseURL: 'http://localhost:8000', // 基础URL使用相对路径，由vite的proxy配置处理
  timeout: 15000, // 请求超时时间
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
request.interceptors.request.use(
  (config:any) => {

    // 从localStorage获取token
    const token = localStorage.getItem('token');
    if (token) {
      config.headers = {
        ...config.headers,
        Authorization: `Bearer ${token}`,
      };
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
request.interceptors.response.use(
  (response: AxiosResponse) => {
    const { data } = response;
    // 这里可以根据后端返回的状态码做相应的处理
    if (data.code===0||data.code) {
      return data;
    }
    // 处理token过期等情况
    if (data.code === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
      return Promise.reject(new Error('登录已过期，请重新登录'));
    }
    return Promise.reject(new Error(data.message || '请求失败'));
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default request;