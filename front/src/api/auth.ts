import request from './request';

// 验证token有效性
export const checkTokenValidity = async () => {
  try {
    const response = await request.get('/gateway/jwt_info');
    // console.log('Token验证API响应:', response);
    return response.code === 0;
  } catch (error) {
    console.error('Token验证失败:', error);
    return false;
  }
};