import * as types from './mutations_types';

export default {

  [types.MUTATION_SET_TOKEN](state, token) {
    state.token = token;
  },

  [types.MUTATION_REMOVE_TOKEN](state) {
    state.token = '';
  },

  [types.MUTATION_USRS](state, users) {
    state.users = users;
  },
};
