import { createApp } from 'vue';
import { Quasar, Notify } from 'quasar';

import '@quasar/extras/material-icons/material-icons.css';
import 'quasar/src/css/index.sass';
import App from './App.vue';

import router from './router';
import { configureClient } from 'api-client';
import { createPinia } from 'pinia';

configureClient({
  protocol: import.meta.env.api_protocol as string,
  host: import.meta.env.api_host as string,
  port: import.meta.env.api_port as string,
});

const pinia = createPinia();
const myApp = createApp(App);

myApp.use(router);
myApp.use(pinia)

myApp.use(Quasar, {
  plugins: {
    Notify
  },
  config: {
    notify: { /* look at QuasarConfOptions from the API card */ }
  }
});

myApp.mount('#app');
