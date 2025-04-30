import { useState, useEffect } from 'react';

/**
 * 自定义Hook，用于监听窗口高度变化
 * @returns {number} 当前窗口高度
 */
const useWindowHeight = (): number => {
  const [windowHeight, setWindowHeight] = useState<number>(0);

  useEffect(() => {
    const handleResize = () => {
      setWindowHeight(window.innerHeight);
    };
    
    // 初始化时立即获取窗口高度
    handleResize();
    
    // 添加事件监听器
    window.addEventListener('resize', handleResize);

    // 清理函数
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []); // 空依赖数组，只在组件挂载和卸载时执行

  return windowHeight;
};

export default useWindowHeight;