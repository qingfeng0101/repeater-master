import type { RouteRecordRaw } from 'vue-router'

import DefaultLayout from '@/layout/DefaultLayout.vue'

export const user: RouteRecordRaw = {
    path: '/user',
    name: 'User',
    component: DefaultLayout,
    meta: {
        hideInMenu: true,
    },
    children: [
        {
            path: 'password',
            name: 'UserChangePassword',
            component: () => import('@/views/user/change-password/index.vue'),
        },
    ]
}