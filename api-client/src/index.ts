import axios, { Axios } from 'axios';
import { 
  ApiError, 
  UsersSignUpDraft, 
  UsersSignInDraft,
  UsersToken, 
  UserPublic 
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
