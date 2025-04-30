import React from 'react';
import { Card, Button, Space, message } from 'antd';
import { HeartOutlined, HeartFilled, SoundOutlined } from '@ant-design/icons';
import './style.css';
import useWindowHeight from '../../hooks/useWindowHeight';

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