import axios from '@/utils/axios';
import store from '@/utils/store';
import * as types from './mutations_types';

export default {

  // 登陆操作
  login({
    commit,
  }, payload) {
    return axios.post('/login', payload)
      .then((resp) => {
        // 保存 token 到 localStorage
        store.set('authToken', resp.data.token);
        // 设置 state.token
        commit(types.MUTATION_SET_TOKEN, resp.data.token);
        return resp;
      });
  },

  // 登出操作
  logout: ({
    commit,
  }) => {
    // 删除 localStorage 内的 authToken
    store.remove('authToken');
    // 清空 state.token
    commit(types.MUTATION_REMOVE_TOKEN);
  },

  get_all_users: ({
    commit,
  }) => axios.get('/user/all')
    .then((resp) => {
      commit(types.MUTATION_USRS, resp.data.users);
      return resp;
    }),

  create_user: ({
    dispatch,
  }, user) => axios.post('/user/', user)
    .then((resp) => {
      dispatch('get_all_users');
      return resp;
    }),
};
