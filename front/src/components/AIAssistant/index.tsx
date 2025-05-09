import React, { useState, useEffect, useRef } from 'react';
import { Card, Button, Input, List, Typography, Spin, message, Statistic, Row, Col, Skeleton, Tabs, Tag, Empty } from 'antd';
import { SendOutlined, RobotOutlined, LineChartOutlined, SearchOutlined, BookOutlined } from '@ant-design/icons';
import { analyzeLearningRecord, generateLearningPlan, formatLearningPlan, getKnowledgeList, searchKnowledge } from './utils';
import './style.css';
import useWindowHeight from '../../hooks/useWindowHeight';
import InfiniteScroll from 'react-infinite-scroll-component';

const { TextArea } = Input;
const { Title, Paragraph } = Typography;

/**
 * 消息接口定义
 * @interface Message
 * @property {string} content - 消息内容
 * @property {'user' | 'ai'} type - 消息类型，用户消息或AI消息
 * @property {number} timestamp - 消息发送时间戳
 */
interface Message {
  content: string;
  type: 'user' | 'ai';
  timestamp: number;
}

/**
 * 知识库条目接口定义
 * @interface KnowledgeItem
 * @property {string} id - 知识条目唯一标识
 * @property {string} title - 知识条目标题
 * @property {string} content - 知识条目内容
 * @property {string} category - 知识条目分类
 * @property {string[]} tags - 知识条目标签列表
 * @property {string} createdAt - 知识条目创建时间
 */
interface KnowledgeItem {
  id: string;
  title: string;
  content: string;
  category: string;
  tags: string[];
  createdAt: string;
}

/**
 * AI学习助手组件
 * @component
 * @description 提供智能对话和知识库查询功能的AI助手组件。
 * 包含以下主要功能：
 * 1. 智能对话：与AI进行实时对话，获取学习建议和答疑
 * 2. 知识库查询：搜索和浏览相关学习知识
 * 3. 学习数据统计：展示用户的学习进度和成果
 * 
 * @returns {JSX.Element} 返回渲染的AI助手组件
 */
const AIAssistant: React.FC = () => {
  const [activeTab, setActiveTab] = useState('chat');
  const [messages, setMessages] = useState<Message[]>([]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);
  const [knowledgeList, setKnowledgeList] = useState<KnowledgeItem[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
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
    height: `${windowHeight - 326}px`,
  }
  const heightknowledgeBase={
    height: `${windowHeight - 282}px`,
  }

  const messageContainerRef = useRef<HTMLDivElement>(null);
  // 滚动到最底部
  useEffect(() => {
    if (messageContainerRef.current) {
      messageContainerRef.current.scrollTop = messageContainerRef.current.scrollHeight;
    }
  }, [messages]);
  useEffect(() => {
    const fetchKnowledge = async () => {
      try {
        const knowledge = await getKnowledgeList();
        setKnowledgeList(knowledge);
      } catch (error) {
        console.error('获取知识库失败:', error);
      }
    };
    fetchKnowledge();
  }, []);

  const handleSearch = async () => {
    if (!searchQuery.trim()) {
      const knowledge = await getKnowledgeList();
      setKnowledgeList(knowledge);
      return;
    }
    try {
      const results = await searchKnowledge(searchQuery);
      setKnowledgeList(results);
    } catch (error) {
      console.error('搜索知识库失败:', error);
    }
  };

  const renderKnowledgeBase = () => (
    <div className="knowledge-base-container" style={heightknowledgeBase}>
      <div className="search-container">
        <Input
          placeholder="搜索知识库"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          onPressEnter={handleSearch}
          suffix={<SearchOutlined onClick={handleSearch} style={{ cursor: 'pointer' }} />}
        />
      </div>
      {knowledgeList.length > 0 ? (
        <List
          itemLayout="vertical"
          dataSource={knowledgeList}
          renderItem={(item) => (
            <List.Item
              key={item.id}
              extra={
                <div className="knowledge-tags">
                  {item.tags.map((tag) => (
                    <Tag key={tag}>{tag}</Tag>
                  ))}
                </div>
              }
            >
              <List.Item.Meta
                title={item.title}
                description={`分类：${item.category} | 创建时间：${new Date(item.createdAt).toLocaleDateString()}`}
              />
              <div className="knowledge-content">{item.content}</div>
            </List.Item>
          )}
        />
      ) : (
        <Empty description="暂无相关知识" />
      )}
    </div>
  );

  return (
    <div className="ai-assistant-container">
      <Card className="ai-card">
        <Title level={4}>
          <RobotOutlined /> AI学习助手
        </Title>
        <Tabs
          activeKey={activeTab}
          onChange={setActiveTab}
          items={[
            {
              key: 'chat',
              label: (
                <span>
                  <RobotOutlined /> 智能对话
                </span>
              ),
              children: (
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
                ),
              },
              {
                key: 'knowledge',
                label: (
                  <span>
                    <BookOutlined /> 知识库
                  </span>
                ),
                children: renderKnowledgeBase(),
              },
            ]}
          />
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
            style={{ borderRadius: '50%', width: "40px", height: "40px" }}
          >
          </Button>
        </div>
      </Card>
    </div>
  );
};

export default AIAssistant;