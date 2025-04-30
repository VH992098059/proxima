import { createBrowserRouter, Navigate ,Outlet} from 'react-router-dom';
import LoginRegister from '../auth/LoginRegister';
import WordLearning from '../pages/WordLearning';
import Home from '../pages/Home';
import Profile from '../pages/Profile';
import MainLayout from '../layouts/MainLayout';
import { MyPerson } from '../pages/MyPerson';
import AIAssistant from '../components/AIAssistant';
import ListeningExercise from '../components/ListeningExercise';
import { message } from 'antd';
// 检查登录状态
const isAuthenticated = () => {
  const token = localStorage.getItem('token');
  return !!token;
};
// 需要登录的路由守卫
const RequireAuth = ({ children }: { children: TSX.Element }) => {
  return isAuthenticated() ? 
  children : 
  <>
    {message.error('请先登录')}
    <Navigate to="/login" />
  </>;
};

const router = createBrowserRouter([
  {
    path: '/',
    element: <MainLayout><Outlet /></MainLayout>,
    children: [
      { index: true, element: <Home /> },
    ]
  },
  {
    path: '/login',
    element: <LoginRegister />
  },
  {
    path: '/word-learning',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <WordLearning /> },
    ]
  },
  {
    path: '/profile',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <Profile /> },
    ]
  },
  {
    path: '/mypersonal',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <MyPerson /> },
    ]
  },
  {
    path: '/ai-assistant',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <AIAssistant /> },
    ]
  },
  {
    path: '/listening-exercise',
    element: <RequireAuth><MainLayout><Outlet /></MainLayout></RequireAuth>,
    children: [
      { index: true, element: <ListeningExercise /> },
    ]
  }
]);


export default router;