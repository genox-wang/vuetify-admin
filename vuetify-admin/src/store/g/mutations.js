import colors from 'vuetify/es5/util/colors';
import store from '@/utils/store';
import * as types from './mutations_types';


export default {

  [types.MUTATION_ALERT_ERROR](state, text) {
    state.snackbar = true;
    state.snackbarText = text;
    state.snackbarColor = 'error';
  },

  [types.MUTATION_ALERT_WARN](state, text) {
    state.snackbar = true;
    state.snackbarText = text;
    state.snackbarColor = 'warn';
  },

  [types.MUTATION_ALERT_SUCCESS](state, text) {
    state.snackbar = true;
    state.snackbarText = text;
    state.snackbarColor = 'success';
  },

  [types.MUTATION_ALERT_CLOSE](state) {
    state.snackbar = false;
  },

  [types.MUTATION_G_UPDATE_THEME_COLOR](state, payload) {
    state.themeColor = payload;
    state.vue.$vuetify.theme.primary = colors[payload.field].darken1;
    store.set('theme_color', state.themeColor);
  },

  [types.MUTATION_G_UPDATE_THEME_IS_DARK](state, payload) {
    state.themeIsDark = payload;
    store.set('theme_is_dark', state.themeIsDark);
  },

  [types.MUTATION_G_INIT_THEME_COLOR](state) {
    const themeColor = store.get('theme_color');
    if (themeColor !== null) {
      state.themeColor = themeColor;
    }
    state.vue.$vuetify.theme.primary = colors[state.themeColor.field].darken1;
  },

  [types.MUTATION_G_INIT_THEME_IS_DARK](state) {
    const themeIsDark = store.get('theme_is_dark');
    if (themeIsDark !== null) {
      state.themeIsDark = themeIsDark;
    }
  },

  [types.MUTATION_G_ATTACH_VUE](state, payload) {
    state.vue = payload;
  },
};
