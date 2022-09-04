import { Router } from 'vue-router';

export interface RouterMiddlewareParam {
  router: Router,
  next: Function
};

export type RouterMiddleware = (param: RouterMiddlewareParam) => void;