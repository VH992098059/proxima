import { createBrowserRouter, Navigate, Outlet, useNavigate } from 'react-router-dom';
import { lazy, Suspense } from 'react';
import { message, Spin } from 'antd';
import MainLayout from '../layouts/MainLayout';

// 使用懒加载导入组件
const LoginRegister = lazy(() => import('../auth/LoginRegister'));
const WordLearning = lazy(() => import('../pages/WordLearning'));
const Home = lazy(() => import('../pages/Home'));
const Profile = lazy(() => import('../pages/Profile'));
const MyPerson = lazy(() => import('../pages/MyPerson').then(module => ({ default: module.MyPerson })));
// MyPerson组件使用的是命名导出，所以需要特殊处理
const AIAssistant = lazy(() => import('../components/AIAssistant'));
const ListeningExercise = lazy(() => import('../components/ListeningExercise'));
const NotFound = lazy(() => import('../pages/NotFound'));

// 加载中组件
const LoadingComponent = () => (
  <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
    <Spin size="large"  />
  </div>
);
import { useEffect, useState } from 'react';
import { checkTokenValidity } from '../api/auth';
import { useAuth } from '../auth/AuthContext';

// 检查登录状态和token有效性
const RequireAuth = ({ children }: { children: React.ReactNode }) => {
  const [isValidating, setIsValidating] = useState(true);
  const navigate=useNavigate();
  const { logout: authLogout } = useAuth();
  useEffect(() => {
    const validateToken = async () => {
      const token = localStorage.getItem('token');
      if (!token) {
        message.warning('请先登录');
        navigate('/login');
        return;
      }
      // console.log('开始验证token...');
      const isValid = await checkTokenValidity();
      // console.log('token验证结果:', isValid);
      
      if (!isValid) {
        authLogout(); // 使用AuthContext的logout方法，它会同时清除localStorage和更新状态
        message.error('登录已过期，请重新登录');
        navigate('/login');
        return;
      }
      setIsValidating(false);
    };

    validateToken();
  }, []);

  if (isValidating) {
    return <LoadingComponent />;
  }
  return <>{children}</>;
};

const router = createBrowserRouter([
  {
    path: '/',
    element: <MainLayout><Outlet /></MainLayout>,
    children: [
      { index: true, element: <Suspense fallback={<LoadingComponent />}><Home /></Suspense> },
    ]
  },
  {
    path: '/login',
    element: <Suspense fallback={<LoadingComponent />}><LoginRegister /></Suspense>
  },
  {
    path: '/word-learning',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <Suspense fallback={<LoadingComponent />}><WordLearning /></Suspense> },
    ]
  },
  {
    path: '/profile',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <Suspense fallback={<LoadingComponent />}><Profile /></Suspense> },
    ]
  },
  {
    path: '/mypersonal',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <Suspense fallback={<LoadingComponent />}><MyPerson /></Suspense> },
    ]
  },
  {
    path: '/ai-assistant',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <Suspense fallback={<LoadingComponent />}><AIAssistant /></Suspense> },
    ]
  },
  {
    path: '/listening-exercise',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <Suspense fallback={<LoadingComponent />}><ListeningExercise /></Suspense> },
    ]
  },
  {
    path: '*',
    element: <Suspense fallback={<LoadingComponent />}><NotFound /></Suspense>
  }
]);


export default router;