import axios from '@/utils/axios';
import * as types from './mutations_types';

export default {

  get_all_servers: ({
    commit,
  }) => axios.get('/server/all')
    .then((resp) => {
      commit(types.MUTATION_SERVERS, resp.data.servers);
    }),

  create_server: ({
    dispatch,
  }, server) => axios.post('/server/', server)
    .then((resp) => {
      dispatch('get_all_servers');
      return resp;
    }),

  update_server: ({
    dispatch,
  }, server) => axios.put(`/server/${server.id}`, server)
    .then((resp) => {
      dispatch('get_all_servers');
      return resp;
    }),

  delete_server: ({
    dispatch,
  }, serverId) => axios.delete(`/server/${serverId}`)
    .then((resp) => {
      dispatch('get_all_servers');
      return resp;
    }),
};
