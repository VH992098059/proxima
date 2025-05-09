import React from 'react';
import { Card, Button, Space, message } from 'antd';
import { HeartOutlined, HeartFilled, SoundOutlined } from '@ant-design/icons';
import './style.css';
import useWindowHeight from '../../hooks/useWindowHeight';

/**
 * 单词卡片组件属性接口
 * @interface WordCardProps
 * @property {string} word - 单词
 * @property {string} phonetic - 音标
 * @property {string} meaning - 单词释义
 * @property {string} example - 示例句子
 * @property {string} [partOfSpeech] - 词性
 * @property {string[]} [synonyms] - 同义词列表
 * @property {string[]} [antonyms] - 反义词列表
 * @property {'easy' | 'medium' | 'hard'} [difficulty] - 单词难度等级
 * @property {boolean} [isFavorite] - 是否已收藏
 * @property {() => void} [onFavoriteToggle] - 切换收藏状态的回调函数
 * @property {() => void} [onNext] - 切换到下一个单词的回调函数
 */
interface WordCardProps {
  word: string;
  phonetic: string;
  meaning: string;
  example: string;
  partOfSpeech?: string;
  synonyms?: string[];
  antonyms?: string[];
  difficulty?: 'easy' | 'medium' | 'hard';
  isFavorite?: boolean;
  onFavoriteToggle?: () => void;
  onNext?: () => void;
}

/**
 * 单词卡片组件
 * @component
 * @description 展示单词学习内容的卡片组件。
 * 主要功能包括：
 * 1. 展示单词基本信息：单词、音标、释义、例句
 * 2. 展示单词扩展信息：词性、同义词、反义词
 * 3. 单词难度标识
 * 4. 单词收藏功能
 * 5. 单词发音功能（开发中）
 * 
 * @param {WordCardProps} props - 组件属性
 * @returns {JSX.Element} 返回渲染的单词卡片组件
 */
const WordCard: React.FC<WordCardProps> = ({
  word,
  phonetic,
  meaning,
  example,
  partOfSpeech = '',
  synonyms = [],
  antonyms = [],
  difficulty = 'medium',
  isFavorite = false,
  onFavoriteToggle,
  onNext,
}) => {
  const playPronunciation = () => {
    // TODO: 实现发音功能
    message.info('发音功能开发中');
  };
  const windowHeight=useWindowHeight();
  const heightStyle = {
    height: `${windowHeight - 316}px`,
  }
  return (
    <>
      <Card className="word-card" style={heightStyle}>
        <div className="word-header">
          <h2 className="word-title">{word}</h2>
          <Button
            type="text"
            icon={<SoundOutlined />}
            onClick={playPronunciation}
            className="pronunciation-btn"
          />
        </div>
        <p className="phonetic">{phonetic}</p>
        <div className="word-content">
          <div className="word-info">
            {partOfSpeech && <span className="part-of-speech">{partOfSpeech}</span>}
            <span className={`difficulty-badge ${difficulty}`}>{difficulty}</span>
          </div>
          <p className="meaning">{meaning}</p>
          <p className="example">{example}</p>
          {synonyms.length > 0 && (
            <div className="word-relations">
              <p className="relation-title">同义词：</p>
              <p className="relation-words">{synonyms.join(', ')}</p>
            </div>
          )}
          {antonyms.length > 0 && (
            <div className="word-relations">
              <p className="relation-title">反义词：</p>
              <p className="relation-words">{antonyms.join(', ')}</p>
            </div>
          )}
        </div>
        <div className="word-actions">
          <Space>
            <Button
              type="text"
              icon={isFavorite ? <HeartFilled style={{color:"red"}}/> : <HeartOutlined />}
              onClick={onFavoriteToggle}
              className={`favorite-btn ${isFavorite ? 'favorite' : ''}`}
            />
          </Space>
        </div>
      </Card>
      <div className="word-actions">
        <Space>
          
          <Button type="primary" onClick={onNext}>
            下一个
          </Button>
        </Space>
      </div>
    </>
  );
};

export default WordCard;