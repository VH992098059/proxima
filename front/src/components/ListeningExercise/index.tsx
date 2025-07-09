import React, { useRef, useState } from 'react';
import { Card, Button, Typography, Progress, message, Tag } from 'antd';
import { SoundOutlined, AudioOutlined, CheckOutlined } from '@ant-design/icons';
import './style.css';
import useWindowHeight from '../../hooks/useWindowHeight';

const { Title, Paragraph } = Typography;

/**
 * 听力练习组件属性接口
 * @interface ListeningExerciseProps
 * @property {() => void} [onNext] - 切换到下一个练习的回调函数
 */
interface ListeningExerciseProps {
  onNext?: () => void;
}

/**
 * 听力练习组件
 * @component
 * @description 提供英语听力练习功能的组件。
 * 主要功能包括：
 * 1. 音频播放：播放标准发音
 * 2. 语音录制：录制用户跟读音频
 * 3. 发音评分：对用户发音进行评估和打分
 * 4. 进度显示：显示录音进度和评分结果
 * 
 * @param {ListeningExerciseProps} props - 组件属性
 * @param {() => void} [props.onNext] - 切换到下一个练习的回调函数
 * 
 * @returns {JSX.Element} 返回渲染的听力练习组件
 */
const ListeningExercise: React.FC<ListeningExerciseProps> = ({
  onNext,
}) => {
  const [isPlaying, setIsPlaying] = useState(false);
  const [isRecording, setIsRecording] = useState(false);
  const [progress, setProgress] = useState(0);
  const [score, setScore] = useState<number | null>(null);
  const audioRef = useRef<HTMLAudioElement | null>(null);
  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const [audioChunks, setAudioChunks] = useState<Blob[]>([]);
  const [translation,setTranslation]=useState<string>('');
  const [sentence,setSentence]=useState<string>('');
  const playAudio = () => {
    if (!audioRef.current) {
      audioRef.current = new Audio();
      audioRef.current.onended = () => {
        setIsPlaying(false);
      };
    }
    audioRef.current.play();
    setIsPlaying(true);
  };

  const startRecording = async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      const mediaRecorder = new MediaRecorder(stream);
      mediaRecorderRef.current = mediaRecorder;
      setAudioChunks([]);

      mediaRecorder.ondataavailable = (event) => {
        if (event.data.size > 0) {
          setAudioChunks((chunks) => [...chunks, event.data]);
        }
      };

      mediaRecorder.onstop = () => {
        const audioBlob = new Blob(audioChunks, { type: 'audio/wav' });
        // 这里可以调用后端API进行发音评估
        const mockScore = Math.floor(Math.random() * 20) + 80; // 模拟评分
        setScore(mockScore);
        setProgress(100);
      };

      mediaRecorder.start();
      setIsRecording(true);
      setProgress(0);
      setScore(null);

      // 3秒后自动停止录音
      setTimeout(() => {
        if (mediaRecorderRef.current && mediaRecorderRef.current.state === 'recording') {
          mediaRecorderRef.current.stop();
          setIsRecording(false);
          stream.getTracks().forEach(track => track.stop());
        }
      }, 3000);
    } catch (error) {
      message.error('无法访问麦克风');
      console.error('录音错误:', error);
    }
  };
  const windowHeight = useWindowHeight();
  const heightStyle = {
    height: `${windowHeight - 450}px`,
  }
  return (
    <Card className="listening-exercise-card">
      <Title level={4}>听力练习</Title>
      <div className="exercise-content" style={heightStyle}>
        <div style={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          height: '100%',
          color: '#999',
          fontSize: '18px'
        }}>
          <Title level={3} style={{ color: '#999', marginBottom: '16px' }}>🚧</Title>
          <Paragraph style={{ color: '#999', fontSize: '16px', textAlign: 'center' }}>
            听力练习功能暂未开发
          </Paragraph>
          <Paragraph style={{ color: '#ccc', fontSize: '14px', textAlign: 'center' }}>
            敬请期待后续版本更新
          </Paragraph>
        </div>
      </div>
    </Card>
  );
};

export default ListeningExercise;