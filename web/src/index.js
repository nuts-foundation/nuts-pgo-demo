import {createApp} from 'vue'
import {createRouter, createWebHashHistory} from 'vue-router'
import './style.css'
import App from './App.vue'
import PGOApp from "./pgo/PGOApp.vue";
import Login from './Login.vue'
import Logout from './Logout.vue'
import NotFound from './NotFound.vue'
import Api from './plugins/api'
import Welcome from './pgo/welcome.vue'
import Consent from './pgo/Consent.vue'
import CreateAccount from './CreateAccount.vue'

const routes = [
    {path: '/', component: Login},
    {
        name: 'login',
        path: '/login',
        component: Login
    },
    {
        name: 'create-account',
        path: '/create-account',
        component: CreateAccount
    },
    {
        name: 'logout',
        path: '/logout',
        component: Logout
    },
    {
        path: '/pgo',
        components: {
            default: PGOApp
        },
        children: [
            {
                path: '',
                name: 'pgo.home',
                redirect: '/pgo/welcome'
            },
            {
                path: 'consent',
                name: "pgo.consent",
                component: Consent
            },
            {
                path: 'welcome',
                name: 'pgo.welcome',
                component: Welcome
            }
        ],
        meta: {requiresAuth: true}
    },
    {path: '/:pathMatch*', name: 'NotFound', component: NotFound}
]

const router = createRouter({
    // We are using the hash history for simplicity here.
    history: createWebHashHistory(),
    routes // short for `routes: routes`
})

router.beforeEach((to, from) => {
    if (to.meta.requiresAuth) {
        if (localStorage.getItem('session')) {
            return true
        }
        return '/login'
    }
})

const app = createApp(App)

app.use(router)
app.use(Api, {forbiddenRoute: {name: 'logout'}})
app.mount('#app')
