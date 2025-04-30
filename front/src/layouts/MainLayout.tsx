import React from 'react';
import { Button } from 'antd';
import { useNavigate, useLocation,Outlet } from 'react-router-dom';
import { HomeOutlined, BookOutlined, HistoryOutlined, UserOutlined, RobotOutlined, AudioOutlined } from '@ant-design/icons';
import { CSSTransition, TransitionGroup } from 'react-transition-group';
import './style.css';
import '../styles/transitions.css';

interface MainLayoutProps {
  children: React.ReactNode;
}

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  const navigate = useNavigate();
  const location = useLocation();

  const isActive = (path: string) => location.pathname === path;

  return (
    <div className="main-layout">
      {/* <TransitionGroup>
        <div className="main-content">{children}</div>
        
      </TransitionGroup> */}
      <div className="main-content"><Outlet /></div>
      <div className="navigation-bar">
        <Button
          type="text"
          className={`nav-button ${isActive('/') ? 'active' : ''}`}
          onClick={() => navigate('/')}
        >
          <HomeOutlined />
          <span>首页</span>
        </Button>
        {/* <Button
          type="text"
          className={`nav-button ${isActive('/word-learning') ? 'active' : ''}`}
          onClick={() => navigate('/word-learning')}
        >
          <BookOutlined />
          <span>学习</span>
        </Button> */}
      
        <Button
          type="text"
          className={`nav-button ${isActive('/ai-assistant') ? 'active' : ''}`}
          onClick={() => navigate('/ai-assistant')}
        >
          <RobotOutlined />
          <span>AI助手</span>
        </Button>
        <Button
          type="text"
          className={`nav-button ${isActive('/listening-exercise') ? 'active' : ''}`}
          onClick={() => navigate('/listening-exercise')}
        >
          <AudioOutlined />
          <span>听力练习</span>
        </Button>
        <Button
          type="text"
          className={`nav-button ${isActive('/mypersonal') ? 'active' : ''}`}
          onClick={() => navigate('/mypersonal')}
        >
          <UserOutlined />
          <span>我的</span>
        </Button>
      </div>
    </div>
  );
};

export default MainLayout;