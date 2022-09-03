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
}