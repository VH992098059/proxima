import { Avatar, Button, Card, message, Col, Row } from "antd";
import "./style.css"
import Icon, { HeartFilled, LogoutOutlined, UserOutlined } from "@ant-design/icons";
import { Link, useNavigate } from 'react-router-dom';
export const MyPerson: React.FC = () => {
  const navigate = useNavigate();
  const handleLogout = () => {
    localStorage.removeItem('token');
    message.success('退出登录成功');
    navigate('/');
  };
  return (
    <div className="myperson-main">
      <Card className="myperson-card">

        <div className="myperson-flex-row">
          <div style={{ display: "flex", flexDirection: "column", width: "100%", height: "85vh", justifyContent: "space-between" }}>
            <Row gutter={[8, 8]}>
              <Col span={24}>
                <div className="myperson-item">
                  <Avatar size={80} icon={<UserOutlined />} />
                  <span style={{ marginLeft: 10 }} >用户名：lingbe32131231231</span>
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
