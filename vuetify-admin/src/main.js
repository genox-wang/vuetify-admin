// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue';
import Vuetify from 'vuetify';
import axios from '@/utils/axios';
import { MUTATION_G_ATTACH_VUE, MUTATION_G_INIT_THEME_COLOR, MUTATION_G_INIT_THEME_IS_DARK } from '@/store/g/mutations_types';
import lodash from 'lodash';
import 'material-design-icons/iconfont/material-icons.css';
import 'vuetify/dist/vuetify.min.css';
import App from './App';
import router from './router';
import store from './store';


Vue.config.productionTip = false;
Vue.prototype.$http = axios;
Vue.prototype._ = lodash;

Vue.use(Vuetify);


router.beforeEach((to, from, next) => {
  if (to.matched.some(r => r.meta.requireAuth)) { // 判断是否需要认证
    if (store.state.user.token) { // 判断是否已经登录
      next();
    } else {
      next({
        name: 'login',
        query: { redirect: to.fullPath }, // 登录成功后重定向到当前页面
      });
    }
  } else {
    next();
  }
});


/* eslint-disable no-new */
const vue = new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>',
});

// attach vue to vuex
store.commit(MUTATION_G_ATTACH_VUE, vue);
// init vuetify theme and dark or not
store.commit(MUTATION_G_INIT_THEME_COLOR);
store.commit(MUTATION_G_INIT_THEME_IS_DARK);

// axios 拦截请求
axios.interceptors.request.use((config) => {
  const c = config;
  if (store.state.user.token) {
    // 存在 token 添加 Authorization 请求头
    c.headers = { Authorization: store.state.user.token };
  }
  return c;
},
error => Promise.reject(error));

// axios 拦截响应
axios.interceptors.response.use((response) => {
  if (response.data.msg) {
    store.dispatch('alert_success', response.data.msg);
  }
  return response;
},
(error) => {
  if (error.response) {
    store.dispatch('alert_error', `${error.response.status}:${error.response.data.msg}`);
  } else {
    store.dispatch('alert_error', error.message);
  }

  if (error.response) {
  // 401 表示认证失败
    if (error.response.status === 401) {
    // 登陆操作
      store.dispatch('logout');
      // 返回登录页面
      router.push({ name: 'login' });
    }
  }

  return Promise.reject(error);
});
