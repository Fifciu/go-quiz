const routes = [
  { path: '/', component: () => import('./pages/Home.vue') },
  { path: '/user-dashboard', component: () => import('./pages/UserDashboard.vue') }
];

export default routes; 
