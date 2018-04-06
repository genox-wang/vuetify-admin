import Vue from 'vue';
import Vuex from 'vuex';

import user from './user';
import g from './g';

Vue.use(Vuex);

const store = new Vuex.Store({

  modules: {
    user,
    g,
  },

});

// vuex hot reloading
if (module.hot) {
  module.hot.accept([
    './user',
    './g',
  ], () => {
    store.hotUpdate({
      modules: {
        user: require('./user').default, // eslint-disable-line global-require
        g: require('./g').default, // eslint-disable-line global-require
      },
    });
  });
}

export default store;
