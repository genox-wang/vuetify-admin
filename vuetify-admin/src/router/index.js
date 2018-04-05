import Vue from 'vue';
import Router from 'vue-router';
import _ from 'lodash';
import googleLayout from '@/layouts/google_layout';
import loginLayout from '@/layouts/login_layout';

const items = require('./nav.json');

Vue.use(Router);


// Route helper function for lazy loading
function route({ path, name, component, meta }) {
  return {
    path,
    name,
    meta,
    component: () => import(
      /* webpackChunkName: "routes" */
      /* webpackMode: "lazy-once" */
      `../pages/${component}.vue`,
    ),
  };
}

const routes = [
  {
    path: '/app',
    component: googleLayout,
    children: [],
  },
  {
    path: '/',
    redirect: { name: 'Dashboard' },
  },
  {
    path: '/login',
    name: 'login',
    component: loginLayout,
  },
];

if (_.isArray(items)) {
  _.forEach(items, (item) => {
    if (item.path) {
      routes[0].children.push(route(item));
    }
    if (_.isArray(item.items)) {
      _.forEach(item.items, (nav) => {
        routes[0].children.push(route(nav));
      });
    }
  });
}


const router = new Router({
  async scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    }

    if (to.hash) {
      return {
        selector: to.hash,
        offset: { y: 80 },
      };
    }

    return { y: 0 };
  },
  routes,
});

export default router;
