
import './App.css'
import { RouterProvider } from 'react-router-dom';
import router from './router';
import { AuthProvider } from './auth/authcontext'; // 导入 AuthProvider

function App() {
  return (
    <AuthProvider> {/* 使用 AuthProvider 包裹应用 */}
      <RouterProvider router={router} />
    </AuthProvider>
  );
}

export default App
