import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import loggedUser from "./global";
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import PhotoCard from "./components/PhotoCard.vue";



import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$loggedUser = loggedUser;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("PhotoCard", PhotoCard);
app.use(router)
app.mount('#app')



