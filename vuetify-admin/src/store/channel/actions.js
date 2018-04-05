import axios from '@/utils/axios';
import * as types from './mutations_types';

export default {

  get_all_channels: ({
    commit,
  }) => axios.get('/channel/all')
    .then((resp) => {
      commit(types.MUTATION_CHANNELS, resp.data.channels);
    }),

  create_channel: ({
    dispatch,
  }, channel) => axios.post('/channel/', channel)
    .then((resp) => {
      dispatch('get_all_channels');
      return resp;
    }),

  update_channel: ({
    dispatch,
  }, channel) => axios.put(`/channel/${channel.id}`, channel)
    .then((resp) => {
      dispatch('get_all_channels');
      return resp;
    }),

  delete_channel: ({
    dispatch,
  }, channelId) => axios.delete(`/channel/${channelId}`)
    .then((resp) => {
      dispatch('get_all_channels');
      return resp;
    }),
};
