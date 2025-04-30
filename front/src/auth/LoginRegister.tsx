import React, { useState } from 'react';
import { Tabs, Form, Input, Button, message, Card } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';
import { Navigate, replace, useNavigate } from 'react-router-dom';

const { TabPane } = Tabs;

const LoginRegister: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const onFinish = (values: any, type: 'login' | 'register') => {
    setLoading(true);
    setTimeout(() => {
      setLoading(false);
      if (type === 'login') {
        localStorage.setItem('token', 'user-token');
        message.success('登录成功');
        navigate('/',{replace:true})
      } else {
        message.success('注册成功');
      }
    }, 1000);
  };

  return (
    <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', background: '#f5f5f5',height:"100vh" }}>
      <Button type="primary" style={{position:"absolute",zIndex:"1000",top:"25px",left:"25px"}} onClick={()=>{navigate("/")}}>返回</Button>

      <Card style={{ width: 350 }}>
        <Tabs defaultActiveKey="login" centered>
          <TabPane tab="登录" key="login">
            <Form name="login" onFinish={v => onFinish(v, 'login')}>
              <Form.Item name="username" rules={[{ required: true, message: '请输入用户名' }]}>  
                <Input prefix={<UserOutlined />} placeholder="用户名" />
              </Form.Item>
              <Form.Item name="password" rules={[{ required: true, message: '请输入密码' }]}>  
                <Input.Password prefix={<LockOutlined />} placeholder="密码" />
              </Form.Item>
              <Form.Item>
                <Button type="primary" htmlType="submit" block loading={loading}>登录</Button>
              </Form.Item>
            </Form>
          </TabPane>
          <TabPane tab="注册" key="register">
            <Form name="register" onFinish={v => onFinish(v, 'register')}>
              <Form.Item name="username" rules={[{ required: true, message: '请输入用户名' }]}>  
                <Input prefix={<UserOutlined />} placeholder="用户名" />
              </Form.Item>
              <Form.Item name="password" rules={[{ required: true, message: '请输入密码' }]}>  
                <Input.Password prefix={<LockOutlined />} placeholder="密码" />
              </Form.Item>
              <Form.Item name="confirm" dependencies={["password"]} rules={[
                { required: true, message: '请确认密码' },
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