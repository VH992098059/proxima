import React, { useState, useEffect } from 'react';
import { Form, Input, Button, Card, Avatar, message, Modal } from 'antd';
import { UserOutlined, LogoutOutlined, LockOutlined, PhoneOutlined, MailOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import './style.css';

interface UserInfo {
  username: string;
  account: string;
  phone: string;
  email: string;
}

interface PasswordForm {
  oldPassword: string;
  newPassword: string;
  confirmPassword: string;
}

const Profile: React.FC = () => {
  const [userInfo, setUserInfo] = useState<UserInfo | null>(null);
  const [editing, setEditing] = useState(false);
  const [isPasswordModalVisible, setIsPasswordModalVisible] = useState(false);
  const [form] = Form.useForm();
  const [passwordForm] = Form.useForm();
  const navigate = useNavigate();

  useEffect(() => {
    // 模拟从后端获取用户信息
    const mockUserInfo = {
      username: '测试用户',
      account: 'test123',
      phone: '13800138000',
      email: 'test@example.com'
    };
    setUserInfo(mockUserInfo);
    form.setFieldsValue(mockUserInfo);
  }, []);

  const handleEdit = () => {
    setEditing(true);
  };

  const handleSave = async (values: UserInfo) => {
    try {
      // TODO: 实现与后端的数据同步
      setUserInfo(values);
      setEditing(false);
      message.success('个人信息更新成功');
    } catch (error) {
      message.error('更新失败，请重试');
    }
  };

  const handleCancel = () => {
    form.setFieldsValue(userInfo);
    setEditing(false);
  };

  const handlePasswordChange = async (values: PasswordForm) => {
    try {
      // TODO: 实现密码修改逻辑
      message.success('密码修改成功');
      setIsPasswordModalVisible(false);
      passwordForm.resetFields();
    } catch (error) {
      message.error('密码修改失败，请重试');
    }
  };

  

  if (!userInfo) {
    return <div>加载中...</div>;
  }

  return (
    <div className="profile-container">
      <Button type="primary" style={{position:"absolute",zIndex:"1000",top:"10px",left:"10px"}} onClick={()=>{navigate(-1)}}>返回</Button>
      <Card title="个人信息" className="profile-card">
        <Form
          form={form}
          layout="vertical"
          onFinish={handleSave}
          initialValues={userInfo}
          className="profile-form"
        //   disabled={!editing}
        >
          <Form.Item
            name="avatar"
            valuePropName="file">
                <div className="profile-header">
                    <Avatar size={75} icon={<UserOutlined />} />
                </div>
            </Form.Item>
          <Form.Item
            label="用户名"
            name="username"
            rules={[{ required: true, message: '请输入用户名' }]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            label="账号"
            name="account"
            rules={[{ required: true, message: '请输入账号' }]}
          >
            <Input prefix={<UserOutlined />} disabled />
          </Form.Item>

          <Form.Item
            label="手机号"
            name="phone"
            rules={[{ required: true, message: '请输入手机号' }, { pattern: /^1[3-9]\d{9}$/, message: '请输入有效的手机号' }]}
          >
            <Input prefix={<PhoneOutlined />} />
          </Form.Item>

          <Form.Item
            label="邮箱"
            name="email"
            rules={[{ required: true, message: '请输入邮箱' }, { type: 'email', message: '请输入有效的邮箱地址' }]}
          >
            <Input prefix={<MailOutlined />} />
          </Form.Item>

          <Form.Item>
            <div className="button-group">
              {editing ? (
                <>
                  <Button type="primary" htmlType="submit">保存</Button>
                  <Button onClick={handleCancel}>取消</Button>
                </>
              ) : (
                <>
                    <Button type="primary" onClick={handleEdit}>编辑</Button>
                    <Button type="default" onClick={() => setIsPasswordModalVisible(true)} icon={<LockOutlined />}>修改密码</Button>
                    {/* <Button  danger onClick={handleLogout} icon={<LogoutOutlined />}>退出登录</Button> */}
                </>
              )}
            </div>
          </Form.Item>
        </Form>

        <Modal
          title="修改密码"
          open={isPasswordModalVisible}
          onCancel={() => {
            setIsPasswordModalVisible(false);
            passwordForm.resetFields();
          }}
          footer={null}
        >
          <Form
            form={passwordForm}
            layout="vertical"
            onFinish={handlePasswordChange}
          >
            <Form.Item
              name="oldPassword"
              label="原密码"
              rules={[{ required: true, message: '请输入原密码' }]}
            >
              <Input.Password prefix={<LockOutlined />} />
            </Form.Item>
            <Form.Item
              name="newPassword"
              label="新密码"
              rules={[{ required: true, message: '请输入新密码' }]}
            >
              <Input.Password prefix={<LockOutlined />} />
            </Form.Item>
            <Form.Item
              name="confirmPassword"
              label="确认新密码"
              dependencies={['newPassword']}
              rules={[
                { required: true, message: '请确认新密码' },
                ({ getFieldValue }) => ({
                  validator(_, value) {
                    if (!value || getFieldValue('newPassword') === value) {
                      return Promise.resolve();
                    }
                    return Promise.reject(new Error('两次输入的密码不一致'));
                  },
                }),
              ]}
            >
              <Input.Password prefix={<LockOutlined />} />
            </Form.Item>
            <Form.Item>
              <Button type="primary" htmlType="submit" block>确认修改</Button>
            </Form.Item>
          </Form>
        </Modal>
      </Card>
    </div>
  );
};

export default Profile;