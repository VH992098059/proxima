import request from "../request";

export interface AiModelParams {
    userId: string;
    message: string;
}

/**
 * 流式AI对话请求
 * @param {AiModelParams} params 模型内容参数
 * @param {(message: string) => void} onMessage 消息回调函数
 * @param {() => void} onError 错误回调函数
 * @param {() => void} onComplete 完成回调函数
 * @returns {() => void} 返回取消订阅的函数
 */
export const getAiModelContentStream = (params: AiModelParams, onMessage: (message: string) => void, onError?: () => void, onComplete?: () => void): () => void => {
    const eventSource = new EventSource(`http://localhost:8000/gateway/ai/chat?userId=${params.userId}&message=${encodeURIComponent(params.message)}`);

    eventSource.onmessage = (event) => {
        onMessage(event.data);
    };

    eventSource.onerror = () => {
        eventSource.close();
        onError?.();
    };

    eventSource.addEventListener('complete', () => {
        eventSource.close();
        onComplete?.();
    });

    return () => {
        eventSource.close();
    };
};