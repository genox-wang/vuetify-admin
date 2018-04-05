import * as types from './mutations_types';

export default {

  [types.MUTATION_CHANNELS](state, channels) {
    state.channels = channels;
  },
};
