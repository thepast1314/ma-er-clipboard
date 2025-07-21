import {createApp} from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import './style.css';

var app = createApp(App);
app.use(createPinia())
app.mount('#app')
