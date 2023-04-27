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