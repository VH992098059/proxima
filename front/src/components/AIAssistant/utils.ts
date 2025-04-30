import { message } from 'antd';

interface LearningRecord {
  wordCount: number;
  studyTime: number;
  correctRate: number;
  lastReviewDate: string;
}

interface LearningPlan {
  dailyWordGoal: number;
  recommendedTime: number;
  focusAreas: string[];
  reviewSchedule: string;
}

// 分析学习记录
export const analyzeLearningRecord = async (userId: string): Promise<LearningRecord> => {
  try {
    // TODO: 从后端获取学习记录数据
    // 模拟数据
    return {
      wordCount: 500,
      studyTime: 120,
      correctRate: 0.85,
      lastReviewDate: new Date().toISOString(),
    };
  } catch (error) {
    message.error('获取学习记录失败');
    throw error;
  }
};

// 生成个性化学习计划
export const generateLearningPlan = async (record: LearningRecord): Promise<LearningPlan> => {
  try {
    // TODO: 调用OpenAI API生成个性化计划
    // 模拟数据
    return {
      dailyWordGoal: Math.round(record.correctRate * 50),
      recommendedTime: Math.max(30, Math.round(record.studyTime * 0.8)),
      focusAreas: ['听力练习', '单词复习', '口语练习'],
      reviewSchedule: '每天复习昨天学习的单词，每周末进行总复习',
    };
  } catch (error) {
    message.error('生成学习计划失败');
    throw error;
  }
};

// 格式化学习计划为可读文本
export const formatLearningPlan = (plan: LearningPlan): string => {
  return `根据你的学习情况，我为你制定了以下学习计划：

1. 每日单词目标：${plan.dailyWordGoal}个
2. 建议学习时间：${plan.recommendedTime}分钟
3. 重点关注领域：${plan.focusAreas.join('、')}
4. 复习安排：${plan.reviewSchedule}`;
};