import request from '../request';

/**
 * 用户登录接口参数类型
 */
export interface LoginParams {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: any;
  code: number;
  message: string;
  data: {
    token: string;
    user: {
      id: number;
      username: string;
      email?: string;
    };
  };
}

/**
 * 用户注册接口参数类型
 */
export interface RegisterParams {
  username: string;
  password: string;
  email?: string;
}

export interface RegisterResponse {
  id: any;
  code: number;
  message: string;
  data: {
    id: number;
    username: string;
  };
}

/**
 * 用户登录请求
 * @param {LoginParams} params - 登录参数
 * @param {string} params.username - 用户名
 * @param {string} params.password - 密码
 * @returns {Promise<LoginResponse>} 返回登录响应，包含用户token和基本信息
 */
export const login = (params: LoginParams) => {
  return request.post<LoginResponse>('/gateway/users/login', params);
};

/**
 * 用户注册请求
 * @param {RegisterParams} params - 注册参数
 * @param {string} params.username - 用户名
 * @param {string} params.password - 密码
 * @param {string} [params.email] - 邮箱（可选）
 * @returns {Promise<RegisterResponse>} 返回注册响应，包含用户ID和用户名
 */
export const register = (params: RegisterParams) => {
  return request.post<RegisterResponse>('/gateway/users/register', params);
};

/**
 * 用户退出登录
 * @returns {Promise<void>} 退出登录的响应
 */
export const logout = () => {
  return request.post('/gateway/users/logout');
};

/**
 * 获取当前登录用户的信息
 * @returns {Promise<any>} 返回用户信息
 */
export const getUserInfo = () => {
  return request.get('/gateway/msg/user/info');
};