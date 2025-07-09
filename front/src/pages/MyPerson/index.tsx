import { Avatar, Button, Card, message, Col, Row } from "antd";
import "./style.css"
import { HeartFilled, LogoutOutlined, UserOutlined } from "@ant-design/icons";
import { Link, useNavigate } from 'react-router-dom';
import { getUserInfo, logout, RegisterParams } from "../../api/login";
import { useAuth } from "../../auth/AuthContext";
import { useEffect, useState } from "react";
export const MyPerson: React.FC = () => {
  const [username,setUsername]=useState('')
  const navigate = useNavigate();
  const { logout: authLogout } = useAuth();
  const handleLogout = async() => {
    const response =await logout();
    if(response.data.logout){
      authLogout(); // 使用AuthContext的logout方法，它会同时清除localStorage和更新状态
      message.success('退出登录成功');
      navigate('/');
    }
  };
  
  useEffect(()=>{
    const getUserMessage=async()=>{
    const userInfo=await getUserInfo()
    setUsername(userInfo.data.username)
  }
  getUserMessage()
  })
  
 
  
  return (
    <div className="myperson-main">
      <Card className="myperson-card">

        <div className="myperson-flex-row">
          <div style={{ display: "flex", flexDirection: "column", width: "100%", height: "85vh", justifyContent: "space-between" }}>
            <Row gutter={[8, 8]}>
              <Col span={24}>
                <div className="myperson-item">
                  <Avatar size={80} icon={<UserOutlined />} />
                  <span style={{ marginLeft: 10 }} >用户名：{username}</span>
                </div>
              </Col>

              <Col span={12}>
                <Card>
                  <Link to="/profile" style={{ color: "#000" }}>
                    <div className="myperson-item-btn">
                      <UserOutlined style={{ color: "#2d9aeb", fontSize: "30px" }} />
                      个人信息
                    </div>
                  </Link>
                </Card>
              </Col>
              <Col span={12}>
                <Card>
                  <Link to="/learning" style={{ color: "#000" }}>
                    <div className="myperson-item-btn">
                      <HeartFilled style={{ color: "red", fontSize: "30px" }} />
                      我的收藏
                    </div>
                  </Link>
                </Card>
              </Col>
            </Row>

            <Button type="primary" danger onClick={handleLogout} icon={<LogoutOutlined />} style={{ height: "60px" }}>退出登录</Button>
          </div>

        </div>
      </Card>

    </div>
  );
};
