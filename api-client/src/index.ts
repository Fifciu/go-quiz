import axios, { Axios } from 'axios';
import { 
  UsersSignUpDraft, 
  UsersSignInDraft,
  UserPublic,
  TestPublic,
  UserResult,
  QuestionWithAnswers,
  UserAnswer,
  FinalUserResult
} from './types';

let instance: Axios;

export const client = {
  users: {
    async signUp (payload: UsersSignUpDraft): Promise<void> {
      return await instance.post('/users/register', payload)
    },
    async signIn (payload: UsersSignInDraft): Promise<void> {
      return await instance.post('/users/login', payload)
    },
    async refreshToken (): Promise<void> {
      return await instance.post('/users/refresh')
    },
    async me (): Promise<UserPublic> {
      return (await instance.post('/users/me')).data;
    }
  },
  tests: {
    async getAll (): Promise<TestPublic[]> {
      return (await instance.get('/tests')).data
    },
    async questionsAndAnswers (testId: number): Promise<QuestionWithAnswers[]> {
      return (await instance.get(`/tests/${testId}/questions/answers`)).data
    },
    async results (): Promise<TestPublic[]> {
      return (await instance.get('/tests/results')).data
    }
  },
  results: {
    async start (testId: number): Promise<UserResult> {
      return (await instance.post(`/results/${testId}/start`)).data
    },

    async finish (resultId: number): Promise<FinalUserResult[]> {
      return (await instance.post(`/results/${resultId}/finish`)).data
    },
  },
  answers: {
    async do (answerId: number): Promise<UserAnswer> {
      return (await instance.put(`/answers/${answerId}`)).data
    }
  }
};

export const configureClient = ({
  protocol,
  host,
  port
}: {
  protocol: string,
  host: string,
  port: string
}) => {
  instance = axios.create({
    baseURL: `${protocol}://${host}:${port}`,
    withCredentials: true
  });
};

export * from './types';
