import React, { useState } from 'react';
import { Progress, message } from 'antd';
import WordCard from '../../components/WordCard';
import './style.css';

// 模拟单词数据，实际项目中应该从API获取
const mockWords = [
  {
    id: 1,
    word: 'example',
    phonetic: '/ɪɡˈzæmpəl/',
    meaning: 'n. 例子，实例；范例',
    example: 'This is an example of how to use the word.',
    partOfSpeech: 'noun',
    synonyms: ['instance', 'sample', 'illustration'],
    antonyms: [],
    difficulty: 'easy',
  },
  {
    id: 2,
    word: 'vocabulary',
    phonetic: '/vəˈkæbjʊˌleri/',
    meaning: 'n. 词汇；词汇量；词表',
    example: 'Reading helps to build your vocabulary.',
    partOfSpeech: 'noun',
    synonyms: ['lexicon', 'terminology', 'words'],
    antonyms: [],
    difficulty: 'medium',
  },
  {
    id: 3,
    word: 'difficult',
    phonetic: '/ˈdɪfɪkəlt/',
    meaning: 'adj. 困难的；难懂的；不易相处的',
    example: 'Learning a new language can be difficult at first.',
    partOfSpeech: 'adjective',
    synonyms: ['hard', 'complicated', 'challenging'],
    antonyms: ['easy', 'simple', 'straightforward'],
    difficulty: 'medium',
  },
];

const WordLearning: React.FC = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [favorites, setFavorites] = useState<number[]>([]);
  const [progress, setProgress] = useState(0);
  const [learningHistory, setLearningHistory] = useState<{
    totalWords: number;
    learnedWords: number;
    lastStudyDate: string;
  }>({ totalWords: mockWords.length, learnedWords: 0, lastStudyDate: new Date().toISOString() });

  const handleNext = () => {
    if (currentIndex < mockWords.length - 1) {
      setCurrentIndex(prev => prev + 1);
      const newProgress = Math.floor(((currentIndex + 2) / mockWords.length) * 100);
      setProgress(newProgress);
      setLearningHistory(prev => ({
        ...prev,
        learnedWords: currentIndex + 2,
        lastStudyDate: new Date().toISOString(),
      }));
    } else {
      message.success('恭喜你完成本组单词学习！');
      // 保存学习记录到本地存储
      localStorage.setItem('wordLearningHistory', JSON.stringify({
        totalWords: mockWords.length,
        learnedWords: mockWords.length,
        lastStudyDate: new Date().toISOString(),
      }));
    }
  };

  const toggleFavorite = () => {
    const currentWordId = mockWords[currentIndex].id;
    setFavorites(prev =>
      prev.includes(currentWordId)
        ? prev.filter(id => id !== currentWordId)
        : [...prev, currentWordId]
    );
  };

  const currentWord = mockWords[currentIndex];

  return (
    <div className="word-learning-container">
      <div className="progress-bar">
        <Progress percent={progress} status="active" />
        <p className="progress-text">
          进度：{currentIndex + 1}/{mockWords.length}
        </p>
      </div>
      <WordCard
        word={currentWord.word}
        phonetic={currentWord.phonetic}
        meaning={currentWord.meaning}
        example={currentWord.example}
        isFavorite={favorites.includes(currentWord.id)}
        onFavoriteToggle={toggleFavorite}
        onNext={handleNext}
      />
    </div>
  );
};

export default WordLearning;