import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import BootstrapVue3 from 'bootstrap-vue-3';

// Import Bootstrap and BootstrapVue CSS files
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue-3/dist/bootstrap-vue-3.css';

const app = createApp(App);

app.use(router);
app.use(store);
app.use(BootstrapVue3);

app.mount('#app');