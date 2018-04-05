import Vue from 'vue';
import Vuex from 'vuex';

import user from './user';
import g from './g';
import channel from './channel';
import server from './server';

Vue.use(Vuex);


const store = new Vuex.Store({

  modules: {
    user,
    g,
    channel,
    server,
  },

});

// vuex hot reloading
if (module.hot) {
  module.hot.accept([
    './user',
    './g',
    './server',
    './channel',
  ], () => {
    store.hotUpdate({
      modules: {
        user: require('./user').default, // eslint-disable-line global-require
        g: require('./g').default, // eslint-disable-line global-require
        channel: require('./channel').default, // eslint-disable-line global-require
        server: require('./server').default, // eslint-disable-line global-require
      },
    });
  });
}

export default store;
