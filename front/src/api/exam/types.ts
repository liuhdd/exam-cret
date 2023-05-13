// 定义 ExamResult 类型
export interface ExamResult {
  exam_id: string;
  student_id: string;
  questions: QuestionResult[];
}

// 定义 QuestionResult 类型
export interface QuestionResult {
  question_id: string;
  answer: string;
  score: number;
}

export interface VerificationResult {
  process: ExamProcess
  ok: boolean
}

export interface ExamProcess {
  exam_id: string;
  student_id: string;
  questions: QuestionInfo[];
}

export interface QuestionInfo {
  question_id: string;
  actions: ActionInfo[];
  score: number;
  scored_by: string;
  scored_time: number;
}

export interface ActionInfo {
  action_id: string;
  answer: string;
  action_time: number;
}
