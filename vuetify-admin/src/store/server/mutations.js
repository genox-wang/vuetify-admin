import * as types from './mutations_types';

export default {

  [types.MUTATION_SERVERS](state, servers) {
    state.servers = servers;
  },
};
