import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  { path: '/', component: () => import('./pages/Home.vue') },
  { path: '/user-dashboard', component: () => import('./pages/UserDashboard.vue') }
];

export default routes; 
