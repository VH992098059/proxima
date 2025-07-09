import React from 'react';
import { Button } from 'antd';
import { useNavigate, useLocation,Outlet } from 'react-router-dom';
import { HomeOutlined, BookOutlined, HistoryOutlined, UserOutlined, RobotOutlined, AudioOutlined } from '@ant-design/icons';
import { CSSTransition, TransitionGroup } from 'react-transition-group';
import './style.css';
import '../styles/transitions.css';

/**
 * MainLayout组件接口定义
 * @interface MainLayoutProps
 * @property {React.ReactNode} children - 子组件，用于渲染主要内容区域
 */
interface MainLayoutProps {
  children: React.ReactNode;
}

/**
 * 主布局组件
 * @component
 * @description 应用程序的主要布局组件，包含导航栏和内容区域。
 * 导航栏包括：首页、AI助手、听力练习和个人中心等功能入口。
 * 
 * @param {MainLayoutProps} props - 组件属性
 * @param {React.ReactNode} props.children - 子组件内容
 * 
 * @returns {JSX.Element} 返回渲染的布局组件
 */
const MainLayout: React.FC<MainLayoutProps> = ({}) => {
  const navigate = useNavigate();
  const location = useLocation();

  const isActive = (path: string) => location.pathname === path;

  return (
    <div className="main-layout">
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