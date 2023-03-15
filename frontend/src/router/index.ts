import { createRouter, createWebHistory } from 'vue-router'
import { useMainStore } from '../store'


import type { RouteRecordRaw } from 'vue-router'
import { LoadingBar } from '@/utils'

import { credential } from './routes/credential'
import { contact } from './routes/contact'
import { prometheus } from './routes/prometheus'
import { sender } from './routes/sender'
import { user } from './routes/user'
import { apikey } from './routes/apikey'

const routes: Array<RouteRecordRaw> = [
    contact,
    sender,
    prometheus,
    credential,
    user,
    apikey,
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/index.vue'),
        props: route => {
            const moduleType = (route.params.module as string) || 'pwd-login'
            return {
                module: moduleType
            }
        },
    },

    {
        path: '/',
        name: 'Home',
        redirect: '/prometheus',
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: () => import('@/views/not-found/index.vue'),
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})
router.beforeEach((to, from, next) => {
    LoadingBar?.start()
    const store = useMainStore()

    if (to.query.access_token) {
        store.token = to.query.access_token as string
    }

    if (to.path !== '/login' && !store.checkAuthToken) {
        next({ name: 'Login', query: { redirect: to.path } })
    } else {
        next()
    }
})
router.afterEach(to => {
    // 设置document title
    document.title = to.meta?.locale || to.name as string
    // 结束 loadingBar
    LoadingBar?.finish()
})
export {
    routes,
    router
}

