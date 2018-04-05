import * as types from './mutations_types';

export default {

  alert_error({
    commit,
  }, payload) {
    commit(types.MUTATION_ALERT_ERROR, payload);
  },

  alert_success({
    commit,
  }, payload) {
    commit(types.MUTATION_ALERT_SUCCESS, payload);
  },

  alert_close({
    commit,
  }) {
    commit(types.MUTATION_ALERT_CLOSE);
  },
};
