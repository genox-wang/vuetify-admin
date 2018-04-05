import store from '@/utils/store';

export default {
  token: store.get('authToken'),
  users: [],
};
