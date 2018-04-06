import { MUTATION_G_SHOW_ALERT, MUTATION_G_CLOSE_ALERT, MUTATION_G_SHOW_NEXT_ALERT } from '@/store/g/mutations_types';

export default {
  store: null,
  use: (store) => {
    this.store = store;
  },

  info: (msg) => {
    this.store.commit(MUTATION_G_SHOW_ALERT, { text: msg, color: 'primary' });
  },

  success: (msg) => {
    this.store.commit(MUTATION_G_SHOW_ALERT, { text: msg, color: 'success' });
  },

  error: (msg) => {
    this.store.commit(MUTATION_G_SHOW_ALERT, { text: msg, color: 'error' });
  },

  warn: (msg) => {
    this.store.commit(MUTATION_G_SHOW_ALERT, { text: msg, color: 'warn' });
  },
  next: () => {
    this.store.commit(MUTATION_G_SHOW_NEXT_ALERT);
  },
  close: () => {
    this.store.commit(MUTATION_G_CLOSE_ALERT);
  },
};
