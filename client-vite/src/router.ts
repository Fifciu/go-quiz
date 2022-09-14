import { RouteRecordRaw } from 'vue-router';
import authMiddleware from './middlewares/auth';
import notAuthMiddleware from './middlewares/notAuth';
import * as VueRouter from 'vue-router';
import { RouterMiddleware } from './types';

const routes: RouteRecordRaw[] = [
  { 
    name: 'home',
    path: '/',
    meta: {
      middleware: notAuthMiddleware
    },
    component: () => import('./pages/Home.vue') 
  },
  {
    name: 'user-dashboard',
    path: '/user-dashboard', 
    meta: {
      middleware: authMiddleware
    }, 
    component: () => import('./pages/UserDashboard.vue')
  },
  {
    name: 'quiz',
    path: '/quiz/:testId', 
    meta: {
      middleware: authMiddleware
    }, 
    component: () => import('./pages/Quiz.vue')
  },
  {
    name: 'quiz-results',
    path: '/quiz/:testId/results', 
    meta: {
      middleware: authMiddleware
    }, 
    component: () => import('./pages/QuizResults.vue')
  }
];

const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes
});

router.beforeEach((to, _, next) => {
  if (to.meta.middleware) {
    return (to.meta.middleware as RouterMiddleware)({ router, next })
  }
  next();
});

export default router; 
