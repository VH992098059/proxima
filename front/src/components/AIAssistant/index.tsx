import React, { useState, useEffect, useRef } from 'react';
import { Card, Button, Input, List, Typography, Spin, message, Statistic, Row, Col, Skeleton } from 'antd';
import { SendOutlined, RobotOutlined, LineChartOutlined } from '@ant-design/icons';
import { analyzeLearningRecord, generateLearningPlan, formatLearningPlan } from './utils';
import './style.css';
import useWindowHeight from '../../hooks/useWindowHeight';
import InfiniteScroll from 'react-infinite-scroll-component';

const { TextArea } = Input;
const { Title, Paragraph } = Typography;

interface Message {
  content: string;
  type: 'user' | 'ai';
  timestamp: number;
}

const AIAssistant: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);
  const [learningStats, setLearningStats] = useState({
    wordCount: 0,
    studyTime: 0,
    correctRate: 0,
  });

  useEffect(() => {
    const fetchLearningRecord = async () => {
      try {
        const record = await analyzeLearningRecord('user123');
        setLearningStats({
          wordCount: record.wordCount,
          studyTime: record.studyTime,
          correctRate: record.correctRate * 100,
        });
      } catch (error) {
        console.error('获取学习记录失败:', error);
      }
    };
    fetchLearningRecord();
  }, []);
  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && e.shiftKey) {

    } else if (e.key === 'Enter') {
      // 阻止默认行为
      e.preventDefault();
      handleSend();
    }
  }
  const handleSend = async () => {
    if (!input.trim()) return;

    const userMessage: Message = {
      content: input,
      type: 'user',
      timestamp: Date.now(),
    };

    setMessages(prev => [...prev, userMessage]);
    setInput('');
    setLoading(true);

    try {
      const record = await analyzeLearningRecord('user123');
      const plan = await generateLearningPlan(record);
      const aiMessage: Message = {
        content: formatLearningPlan(plan),
        type: 'ai',
        timestamp: Date.now(),
      };
      setMessages(prev => [...prev, aiMessage]);
      setLoading(false);
    } catch (error) {
      message.error('获取AI响应失败，请稍后重试');
      setLoading(false);
    }
  };
  const windowHeight = useWindowHeight();
  const heightStyle = {
    height: `${windowHeight - 316}px`,
  }

  const messageContainerRef = useRef<HTMLDivElement>(null);
  // 滚动到最底部
  useEffect(() => {
    if (messageContainerRef.current) {
      messageContainerRef.current.scrollTop = messageContainerRef.current.scrollHeight;
    }
  }, [messages]);
  return (
    <div className="ai-assistant-container">
      <Card className="ai-card">
        <Title level={4}>
          <RobotOutlined /> AI学习助手
        </Title>
        <div className="message-container" style={heightStyle} ref={messageContainerRef}>
          <InfiniteScroll
            dataLength={messages.length}
            next={() => {}}
            hasMore={false}
            scrollableTarget="message-container"
            loader={<Skeleton avatar paragraph={{ rows: 1 }} active />}
          >
            <List
              itemLayout="horizontal"
              dataSource={messages}
              renderItem={msg => (
                <List.Item className={`message ${msg.type}`}>
                  <Paragraph>{msg.content}</Paragraph>
                </List.Item>
              )}

            />
            {loading && (
              <div className="loading-container">
                <Spin />
              </div>
            )}
          </InfiniteScroll>

        </div>
        <div className="input-container">
          <TextArea
            value={input}
            onChange={e => setInput(e.target.value)}
            placeholder="请输入你的问题，例如：帮我分析学习记录并制定计划"
            autoSize={{ minRows: 2, maxRows: 3 }}
            onKeyDown={handleKeyPress}
          />
          <Button
            type="primary"
            icon={<SendOutlined />}
            onClick={handleSend}
            loading={loading}
            style={{ borderRadius: '50%', width: "50px", height: "50px" }}
          >
          </Button>
        </div>
      </Card>
    </div>
  );
};

export default AIAssistant;