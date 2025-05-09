import { message } from 'antd';

interface LearningRecord {
  wordCount: number;
  studyTime: number;
  correctRate: number;
  lastReviewDate: string;
}

interface KnowledgeItem {
  id: string;
  title: string;
  content: string;
  category: string;
  tags: string[];
  createdAt: string;
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

// 获取知识库列表
export const getKnowledgeList = async (): Promise<KnowledgeItem[]> => {
  try {
    // TODO: 从后端获取知识库列表
    // 模拟数据
    return [
      {
        id: '1',
        title: '常用英语口语表达',
        content: '日常生活中最常用的英语口语表达方式和例句...',
        category: '口语',
        tags: ['日常对话', '实用表达'],
        createdAt: new Date().toISOString(),
      },
      {
        id: '2',
        title: '商务英语写作技巧',
        content: '职场中常用的商务英语写作格式和范例...',
        category: '写作',
        tags: ['商务英语', '职场沟通'],
        createdAt: new Date().toISOString(),
      },
    ];
  } catch (error) {
    message.error('获取知识库列表失败');
    throw error;
  }
};

// 搜索知识库
export const searchKnowledge = async (query: string): Promise<KnowledgeItem[]> => {
  try {
    // TODO: 调用后端搜索API
    const allKnowledge = await getKnowledgeList();
    return allKnowledge.filter(item =>
      item.title.toLowerCase().includes(query.toLowerCase()) ||
      item.content.toLowerCase().includes(query.toLowerCase()) ||
      item.tags.some(tag => tag.toLowerCase().includes(query.toLowerCase()))
    );
  } catch (error) {
    message.error('搜索知识库失败');
    throw error;
  }
}
