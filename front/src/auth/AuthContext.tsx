import React, { createContext, useContext, useState, ReactNode } from 'react';

interface AuthContextType {
  isAuthenticated: boolean;
  login: () => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(() => {
    // 初始化时从localStorage读取登录状态
    const storedAuth = localStorage.getItem('token');
    return storedAuth !== null;
  });

  const login = () => {
    // 在实际应用中，这里会处理登录逻辑，例如调用API，成功后设置isAuthenticated为true
    setIsAuthenticated(true);
  };

  const logout = () => {
    // 清除localStorage中的token并更新认证状态
    localStorage.removeItem('token');
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};