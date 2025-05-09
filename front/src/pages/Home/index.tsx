import React from 'react';
import { Card, Row, Col, Button, Progress, Typography, Statistic } from 'antd';
import { BookOutlined, ClockCircleOutlined, RightOutlined} from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import './style.css';

const { Title } = Typography;

const Home: React.FC = () => {
  const navigate = useNavigate();

  const startLearning = () => {
    navigate('/word-learning');
  };

  return (
    <div className="home-container">
      <div className="header-section">
        <Title level={2}>灵比英语本</Title>
        <p className="subtitle">让英语学习更轻松、更高效</p>
      </div>
      <Row gutter={[16, 16]} className="stats-section">
        <Col xs={24} sm={12}>
          <Card>
            <Statistic
              title="已学单词"
              value={150}
              prefix={<BookOutlined />}
              suffix="个"
            />
          </Card>
        </Col>
        <Col xs={24} sm={12}>
          <Card>
            <Statistic
              title="学习时长"
              value={120}
              prefix={<ClockCircleOutlined />}
              suffix="分钟"
            />
          </Card>
        </Col>
      </Row>
      <Card className="progress-card">
        <Title level={4}>今日学习进度</Title>
        <Progress percent={30} status="active" />
        <p className="progress-text">已完成今日目标的30%</p>
      </Card>
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