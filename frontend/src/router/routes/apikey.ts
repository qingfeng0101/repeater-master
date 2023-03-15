import type { RouteRecordRaw } from 'vue-router'

import DefaultLayout from '@/layout/DefaultLayout.vue'

export const apikey: RouteRecordRaw = {
    path: '/apikey',
    name: 'APIKey',
    component: DefaultLayout,
    meta: {
        flatChildrenInMenu: true
    },
    children: [
        {
            path: 'list',
            name: 'APIKeyList',
            component: () => import('@/views/apikey/index.vue'),
            meta: {
                locale: 'APIKeys',
                icon: 'Password'
            },
        } 
    ]
}