import { createApp } from 'vue';
import { Quasar } from 'quasar';

import '@quasar/extras/material-icons/material-icons.css';
import 'quasar/src/css/index.sass';
import App from './App.vue';

import routes from './routes';
import * as VueRouter from 'vue-router';

const router = VueRouter.createRouter({
  history: VueRouter.createWebHashHistory(),
  routes
});

const myApp = createApp(App);

myApp.use(router);

myApp.use(Quasar, {
  plugins: {}
});

myApp.mount('#app');
