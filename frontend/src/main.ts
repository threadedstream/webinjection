import {createApp} from 'vue'
import 'mdb-vue-ui-kit/css/mdb.min.css'
import App from './App.vue'
import axios from "axios"
import router from "./router"
import VueAxios from "vue-axios"

const client = axios.create({
    baseURL: "/api/v1",
})

const app = createApp(App)

app.use(VueAxios, client)
app.use(router)

app.mount("#app")