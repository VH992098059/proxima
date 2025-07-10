import React, { useState } from 'react';
import { Tabs, Form, Input, Button, message, Card } from 'antd';
import { UserOutlined, LockOutlined, SendOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import { login, register, LoginParams, RegisterParams, LoginResponse, RegisterResponse } from '../api/login';
import { useAuth } from './AuthContext';

const { TabPane } = Tabs;
const LoginRegister: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const [registerForm] = Form.useForm();//用于注册表单
  const { login: authLogin } = useAuth();

  const onFinish = async (values: LoginParams | RegisterParams, type: 'login' | 'register') => {
    setLoading(true);
    try {
      if (type === 'login') {
        const response:LoginResponse = await login(values as LoginParams);
        if (response.code === 0 && response.data.token) {
          localStorage.setItem('token', response.data.token);
          authLogin(); // 更新AuthContext的认证状态
          message.success('登录成功');
          navigate('/', { replace: true });
        } else {
          message.warning(response.message || '登录失败');
        }
      } else {
        const response:RegisterResponse = await register(values as RegisterParams);
        if (response.code === 0 && response.data.id) {
          message.success('注册成功');
          registerForm.resetFields();//注册成功后重置表单
          // 切换到登录标签页
          const loginTab = document.querySelector('.ant-tabs-tab[data-node-key="login"]') as HTMLElement;
          if (loginTab) {
            loginTab.click();
          }
        } else {
          message.warning(response.message || '注册失败');
        }
      }
    } catch (error: any) {
      message.error(error.message || error.response?.data?.message || '操作失败，请重试');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', background: '#f5f5f5',height:"100vh" }}>
      <Button type="primary" style={{position:"absolute",zIndex:"1000",top:"25px",left:"25px"}} onClick={()=>{navigate("/")}}>返回</Button>

      <Card style={{ width: 350 }}>
        <Tabs defaultActiveKey="login" centered>
          <TabPane tab="登录" key="login">
            <Form name="login" onFinish={v => onFinish(v, 'login')}>
              <Form.Item 
                name="username" 
                rules={[
                  { required: true, message: '请输入用户名' },
                  { min: 6, message: '用户名长度不能少于6位' }
                ]}
              >  
                <Input prefix={<UserOutlined />} placeholder="用户名" />
              </Form.Item>
              <Form.Item 
                name="password" 
                rules={[
                  { required: true, message: '请输入密码' },
                  { min: 6, message: '密码长度不能少于6位' }
                ]}
              >  
                <Input.Password prefix={<LockOutlined />} placeholder="密码" />
              </Form.Item>
              <Form.Item>
                <Button type="primary" htmlType="submit" block loading={loading}>登录</Button>
              </Form.Item>
            </Form>
          </TabPane>
          <TabPane tab="注册" key="register">
            <Form form={registerForm} name="register" onFinish={v => onFinish(v, 'register')}>
              <Form.Item 
                name="username" 
                rules={[
                  { required: true, message: '请输入用户名' },
                  { min: 6, message: '用户名长度不能少于6位' }
                ]}
              >  
                <Input prefix={<UserOutlined />} placeholder="用户名" />
              </Form.Item>
              <Form.Item 
                name="email" 
                rules={[
                  { required: true, message: '请输入邮箱' },
                  { type: 'email', message: '请输入有效的邮箱地址' }
                ]}
              >  
                <Input prefix={<SendOutlined />} placeholder="邮箱" />
              </Form.Item>
              <Form.Item 
                name="password" 
                rules={[
                  { required: true, message: '请输入密码' },
                  { min: 6, message: '密码长度不能少于6位' }
                ]}
              >  
                <Input.Password prefix={<LockOutlined />} placeholder="密码" />
              </Form.Item>
              <Form.Item 
                name="confirm" 
                dependencies={["password"]} 
                rules={[
                  { required: true, message: '请确认密码' },
                  { min: 6, message: '确认密码长度不能少于6位' },
                  ({ getFieldValue }) => ({
                    validator(_, value) {
                      if (!value || getFieldValue('password') === value) {
                        return Promise.resolve();
                      }
                      return Promise.reject(new Error('两次输入的密码不一致'));
                    },
                  }),
              ]}>
                <Input.Password prefix={<LockOutlined />} placeholder="确认密码" />
              </Form.Item>
              <Form.Item>
                <Button type="primary" htmlType="submit" block loading={loading}>注册</Button>
              </Form.Item>
            </Form>
          </TabPane>
        </Tabs>
      </Card>
    </div>
  );
};

export default LoginRegister;