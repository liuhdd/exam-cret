// 定义 ExamResult 类型
export interface ExamResult {
  exam_id: string;
  student_id: string;
  exam_name: string;
  begin_time: number;
  end_time: number;
  paper: string;
  place: string;
  grade: number;
  questions: QuestionResult[];
}
// 定义 Exam 类型
export interface Exam{
  exam_id: string;
  exam_name: string;
  paper: string;
  paper_id: string;
  place: string;
  begin_time: number;
  end_time: number;
  exam_server: string;
  exam_server_key: string;
  exam_desc: string;
}
// 定义 QuestionResult 类型
export interface QuestionResult {
  question_id: string;
  content: string;
  answer: string;
  score: number;
}

export interface Grade {
  exam_id: string;
  student_id: string;
  student_name: string;
  exam_name: string;
  begin_time: number;
  end_time: number;
  grade: number;
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
