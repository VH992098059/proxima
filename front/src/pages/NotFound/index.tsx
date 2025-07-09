import React from 'react';
import { Button, Result } from 'antd';
import { useNavigate } from 'react-router-dom';
import './style.css';

/**
 * 404页面组件
 * @component
 * @description 当用户访问不存在的页面时显示的404错误页面
 * 提供友好的错误提示和导航选项
 */
const NotFound: React.FC = () => {
  const navigate = useNavigate();

  const handleGoHome = () => {
    navigate('/');
  };

  const handleGoBack = () => {
    navigate(-1);
  };

  return (
    <div className="not-found-container">
      <Result
        status="404"
        title="404"
        subTitle="抱歉，您访问的页面不存在"
        extra={
          <div className="not-found-actions">
            <Button type="primary" onClick={handleGoHome}>
              返回首页
            </Button>
            <Button onClick={handleGoBack}>
              返回上一页
            </Button>
          </div>
        }
      />
      <div className="not-found-tips">
        <h3>🎯 继续您的英语学习之旅</h3>
        <p>不如趁此机会去学习一些新单词吧！</p>
        <Button 
          type="link" 
          onClick={() => navigate('/word-learning')}
          className="learning-link"
        >
          开始单词学习 →
        </Button>
      </div>
    </div>
  );
};

export default NotFound;