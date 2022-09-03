import axios from 'axios';
import { 
  ApiError, 
  UsersSignUpDraft, 
  UsersSignInDraft,
  UsersToken, 
  UserPublic 
} from './types';

const instance = axios.create({
  baseURL: `${process.env.api_protocol}://${process.env.api_host}:${process.env.api_port}`
});

const client = {
  users: {
    async signUp (payload: UsersSignUpDraft): Promise<UsersToken | ApiError> {
      return await instance.post('/users/register', payload)
    },
    async signIn (payload: UsersSignInDraft): Promise<UsersToken | ApiError> {
      return await instance.post('/users/login', payload)
    },
    async refreshToken (): Promise<UsersToken | ApiError> {
      return await instance.post('/users/refresh')
    },
    async me (): Promise<UserPublic | ApiError> {
      return await instance.post('/users/me')
    }
  }
};

export default client;
export * from './types';
