export interface UsersSignUpDraft {
  fullname: string;
  email: string;
  password: string;  
};

export interface UsersSignInDraft {
  email: string;
  password: string;  
};

export interface UsersToken {
  token: string;
  expirationTime: string;
};

export interface UserPublic {
  id: number;
  fullname: string;
  email: string;
};

export interface ApiError {
  message: string;
};

export interface TestPublic {
  id: number;
  title: string;
  description: string;
  image_url: string;
};

export interface UserResult {
  id: number;
  test_id: number;
  user_id: number;
  start_datetime: string;
  finish_datetime?: string;
};

export interface Answer {
  id: number;
  content: string;
};

export interface QuestionWithAnswers {
  id: number;
  content: string;
  answers: Answer[];
};

export interface UserAnswer {
  id: number,
  user_id: number,
  answer_id: number,
  created_at: string
};

export interface FinalUserResult {
  question: string,
  answer: string,
  is_proper: boolean
}