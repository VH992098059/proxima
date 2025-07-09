import React, { useState } from 'react';
import { Progress, message } from 'antd';
import WordCard from '../../components/WordCard';
import './style.css';

/**
 * 模拟单词数据结构
 * @type {Array<{
 *   id: number,
 *   word: string,
 *   phonetic: string,
 *   meaning: string,
 *   example: string,
 *   partOfSpeech: string,
 *   synonyms: string[],
 *   antonyms: string[],
 *   difficulty: 'easy' | 'medium' | 'hard'
 * }>}
 * @description 用于模拟从API获取的单词数据，包含单词的基本信息和学习相关属性
 */
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

/**
 * 单词学习组件
 * @component
 * @description 提供单词学习功能的主要组件。
 * 
 * 主要功能：
 * 1. 单词学习进度追踪：通过Progress组件显示当前学习进度
 * 2. 单词收藏功能：支持将单词添加到收藏列表或从收藏列表中移除
 * 3. 学习历史记录：记录总单词数、已学习单词数和最后学习日期
 * 4. 本地存储：在完成学习后将学习记录保存到localStorage
 * 
 * 状态管理：
 * - currentIndex: 当前学习的单词索引
 * - favorites: 收藏的单词ID列表
 * - progress: 学习进度百分比
 * - learningHistory: 学习历史记录，包含总单词数、已学习数量和最后学习日期
 */
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
    if (currentIndex < mockWords.length - 1) { // 如果当前不是最后一个单词
      const nextShowingIndex = currentIndex + 1; // 这是下一个要显示的单词的索引
      setCurrentIndex(nextShowingIndex);

      // 此时，我们完成了 nextShowingIndex 个单词 (因为索引从0开始)
      const numberOfWordsCompleted = nextShowingIndex;

      const newProgress = Math.floor((numberOfWordsCompleted / mockWords.length) * 100);
      setProgress(newProgress);
      setLearningHistory(prev => ({
        ...prev,
        learnedWords: numberOfWordsCompleted,
        lastStudyDate: new Date().toISOString(),
      }));
    } else { // 这是点击“下一个”并且当前是最后一个单词的情况，意味着所有单词都学习完了
      setProgress(100); // 确保进度是100%
      setLearningHistory(prev => ({
        ...prev,
        learnedWords: mockWords.length, // 确保学习历史记录中的已学习单词数是总数
        lastStudyDate: new Date().toISOString(),
      }));
      message.success('恭喜你完成本组单词学习！');
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