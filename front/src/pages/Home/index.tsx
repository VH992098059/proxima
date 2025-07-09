import React from 'react';
import { Card, Row, Col, Button, Progress, Typography, Statistic } from 'antd';
import { BookOutlined, ClockCircleOutlined, RightOutlined, LoginOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../../auth/AuthContext'; // 导入 useAuth
import './style.css';

const { Title } = Typography;

const Home: React.FC = () => {
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();
  
  
  const startLearning = () => {
    navigate('/word-learning');
  };

  const goToLogin = () => {
    navigate('/login');
  };

  return (
    <div className="home-container">
      <div className="header-section">
        <Title level={2}>灵比英语本</Title>
        <p className="subtitle">让英语学习更轻松、更高效</p>
      </div>

      {isAuthenticated ? (
        <>
          <Row gutter={[16, 16]} className="stats-section">
            <Col xs={24} sm={12}>
              <Card>
                <Statistic
                  title="已学单词"
                  value={150} // 示例数据，后续替换为真实数据
                  prefix={<BookOutlined />}
                  suffix="个"
                />
              </Card>
            </Col>
            <Col xs={24} sm={12}>
              <Card>
                <Statistic
                  title="学习时长"
                  value={120} // 示例数据，后续替换为真实数据
                  prefix={<ClockCircleOutlined />}
                  suffix="分钟"
                />
              </Card>
            </Col>
          </Row>
          <Card className="progress-card">
            <Title level={4}>今日学习进度</Title>
            <Progress percent={30} status="active" /> // 示例数据，后续替换为真实数据
            <p className="progress-text">已完成今日目标的30%</p>
          </Card>
        </>
      ) : (
        <Card className="login-prompt-card">
          <Title level={4}>登录以查看您的学习进度</Title>
          <p>记录您的学习历程，获取个性化学习计划。</p>
          <Button type="primary" icon={<LoginOutlined />} onClick={goToLogin} size="large">
            登录/注册
          </Button>
        </Card>
      )}

      <div className="action-section">
        <Button
          type="primary"
          size="large"
          icon={<RightOutlined />}
          onClick={startLearning}
          className="start-button"
        >
          开始刷词
        </Button>
      </div>
    </div>
  );
};

export default Home;