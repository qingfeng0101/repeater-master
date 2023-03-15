import type { RouteRecordRaw } from 'vue-router'

import DefaultLayout from '@/layout/DefaultLayout.vue'

export const prometheus: RouteRecordRaw = {
    path: '/prometheus',
    name: 'Prometheus',
    redirect: { name: 'PrometheusNotifyRule' },
    component: DefaultLayout,
    meta: {
        locale: 'Prometheus',
        icon: 'Prometheus'
    },
    children: [
        {
            path: 'datasource',
            name: 'PrometheusDatasourceList',
            component: () => import('@/views/prometheus/datasource/index.vue'),
            meta: {
                locale: 'DataSource',
                icon: 'DataSwitching'
            },
        },
        {
            path: 'datasource/add',
            name: 'PrometheusDatasourceAdd',
            component: () => import('@/views/prometheus/datasource/Action.vue'),
            meta: {
                hideInMenu: true
            },
        },
        {
            path: 'datasource/:id/edit',
            name: 'PrometheusDatasourceEdit',
            component: () => import('@/views/prometheus/datasource/Action.vue'),
            props: true,
            meta: {
                hideInMenu: true
            },
        },
        {
            path: 'notify/rule',
            name: 'PrometheusNotifyRule',
            component: () => import('@/views/prometheus/notify-rule/index.vue'),
            meta: {
                locale: 'NotifyRule',
                icon: 'PushRules'
            },
        },
        {
            path: 'notify/rule/cache',
            name: 'PrometheusNotifyRuleCache',
            component: () => import('@/views/prometheus/notify-rule-cache/index.vue'),
            meta: {
                locale: 'NotifyRuleCacheTree',
                icon: 'Tree',
            },
        },
        {
            path: 'notify/rule/add',
            name: 'PrometheusNotifyRuleAdd',
            component: () => import('@/views/prometheus/notify-rule/Action.vue'),
            meta: {
                hideInMenu: true
            },
        },
        {
            path: 'notify/rule/:id/edit',
            name: 'PrometheusNotifyRuleEdit',
            props: true,
            component: () => import('@/views/prometheus/notify-rule/Action.vue'),
            meta: {
                hideInMenu: true,
            },
        },
        {
            path: 'alerts',
            name: 'PrometheusAlerts',
            component: () => import('@/views/prometheus/alerts/index.vue'),
            meta: {
                locale: 'Alerts',
                icon: 'DataSwitching'
            },
        },
        {
            path: 'fakedatasource',
            name: 'FakePrometheusDatasourceList',
            component: () => import('@/views/prometheus/fake-datasource/index.vue'),
            meta: {
                locale: 'FakeDataSource',
                icon: 'DataSwitching'
            },
        },
        {
            path: 'fakedatasource/add',
            name: 'FakePrometheusDatasourceAdd',
            component: () => import('@/views/prometheus/fake-datasource/Action.vue'),
            meta: {
                hideInMenu: true
            },
        },
        {
            path: 'fakedatasource/:id/edit',
            name: 'FakePrometheusDatasourceEdit',
            component: () => import('@/views/prometheus/fake-datasource/Action.vue'),
            props: true,
            meta: {
                hideInMenu: true
            },
        }
    ]
}