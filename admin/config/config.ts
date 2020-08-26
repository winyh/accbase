// https://umijs.org/config/
import { defineConfig } from 'umi';
import defaultSettings from './defaultSettings';
import proxy from './proxy';

const { REACT_APP_ENV } = process.env;

export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  layout: {
    name: 'Accbase',
    locale: true,
    siderWidth: 208,
  },
  locale: {
    // default zh-CN
    default: 'zh-CN',
    // default true, when it is true, will use `navigator.language` overwrite default
    antd: true,
    baseNavigator: true,
  },
  dynamicImport: {
    loading: '@/components/PageLoading/index',
  },
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/docs/routing
  routes: [
    {
      path: '/user',
      layout: false,
      routes: [
        {
          name: 'login',
          path: '/user/login',
          component: './user/login',
        },
      ],
    },
    {
      path: '/welcome',
      name: 'welcome',
      icon: 'smile',
      component: './Welcome',
    },
    {
      path: '/admin',
      name: 'admin',
      icon: 'crown',
      access: 'canAdmin',
      component: './Admin',
      routes: [
        {
          path: '/admin/sub-page',
          name: 'sub-page',
          icon: 'smile',
          component: './Welcome',
        },
      ],
    },
    {
      name: 'list.table-list',
      icon: 'table',
      path: '/list',
      component: './ListTableList',
    },
    {
      name: 'sms',
      icon: 'user',
      path: '/sms',
      routes: [
        {
          name: 'user',
          path: '/sms/user',
          component: './user/list',
        },
        {
          name: 'role',
          path: '/sms/role',
          component: './user/role',
        },
        {
          name: 'power',
          path: '/sms/power',
          component: './user/power',
        },
        {
          name: 'company',
          path: '/sms/company',
          component: './user/company',
        },
        {
          name: 'post',
          path: '/sms/post',
          component: './user/menu',
        },
        {
          name: 'menu',
          path: '/sms/menu',
          component: './user/menu',
        },
      ],
    },
    {
      name: 'system',
      icon: 'appstore',
      path: '/system',
      routes: [
        {
          name: 'info',
          path: '/system/info',
          component: './system/info',
        },
        {
          name: 'log',
          path: '/system/log',
          component: './system/log',
        },
        {
          name: 'dictionary',
          path: '/system/dictionary',
          component: './system/log',
        },
      ],
    },
    {
      path: '/',
      redirect: '/welcome',
    },
    {
      component: './404',
    },
  ],
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    // ...darkTheme,
    'primary-color': defaultSettings.primaryColor,
  },
  // @ts-ignore
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy[REACT_APP_ENV || 'dev'],
  manifest: {
    basePath: '/',
  },
});
