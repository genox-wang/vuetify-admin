<template>
  <v-app id="inspire" :dark="themeIsDark">
    <v-navigation-drawer
      fixed
      :clipped="$vuetify.breakpoint.lgAndUp"
      app
      v-model="drawer"
    >
      <v-list dense>
        <template v-for="item in items">
          <v-list-group
            v-if="item.items"
            v-model="item.model"
            :key="item.name"
            :prepend-icon="item.model ? item.icon : item['icon-alt']"
            append-icon=""
          >
            <v-list-tile slot="activator">
              <v-list-tile-content>
                <v-list-tile-title>
                  {{ item.name }}
                </v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile
              v-model="child.active"
              v-for="(child, i) in item.items"
              :key="i"
              @click="route(child.name)"
              ripple
            >
              <v-list-tile-action v-if="child.icon">
                <v-icon>{{ child.icon }}</v-icon>
              </v-list-tile-action>
              <v-list-tile-content>
                <v-list-tile-title>
                  {{ child.name }}
                </v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
          </v-list-group>
          <v-list-tile v-else :key="item.name" @click="route(item.name)" v-model="item.active">
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>
                {{ item.name }}
              </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </template>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar
      color="primary"
      dark
      app
      :clipped-left="$vuetify.breakpoint.lgAndUp"
      fixed
    >
      <v-toolbar-title style="width: 300px" class="ml-0 pl-3">
        <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
        <span class="hidden-sm-and-down">Vuetify-Admin</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>

      <theme-dark-swith v-if="config.themeEditable"/>

      <theme-color-selection v-if="config.themeEditable"/>

      <v-btn v-if="config.authRequired" icon @click.native="doLogout">
        <v-icon>exit_to_app</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content>
      <router-view></router-view>
    </v-content>
    <base-alert/>
  </v-app>
</template>
<script>
import themeColorSelection from '@/components/theme_color_selection';
import themeDarkSwith from '@/components/theme_dark_switch';
import baseAlert from '@/components/base_alert';
import colors from 'vuetify/es5/util/colors';
import config from '@/config';

const items = require('../router/nav.json');

export default {
  components: {
    themeColorSelection,
    themeDarkSwith,
    baseAlert,
  },

  data: () => ({
    drawer: null,
    items,
    config,
  }),

  computed: {
    themeIsDark() {
      return this.$store.state.g.themeIsDark;
    },
  },

  methods: {
    route(name) {
      this.$router.push({ name });
    },

    doLogout() {
      // this.$http.get('/channel/all');
      this.$store.dispatch('logout');
      this.$router.push({ name: 'login' });
    },

    activeCurrentItem() {
      const curRouteName = this.$route.name;
      if (this._.isArray(items)) {
        items.map((curItem) => {
          const item = curItem;
          item.active = false;
          if (item.path) {
            if (this._.isEqual(item.name, curRouteName)) {
              item.active = true;
            }
          }
          if (this._.isArray(item.items)) {
            item.items.map((curSub) => {
              const sub = curSub;
              sub.active = false;
              if (this._.isEqual(sub.name, curRouteName)) {
                sub.active = true;
              }
              return sub;
            });
          }
          return item;
        });
      }
    },
  },

  watch: {
    themeColor(cur) {
      this.$vuetify.theme.primary = colors[cur.field].darken1;// colors[cur.field].dark1;
    },
    $route: 'activeCurrentItem',
  },

  mounted() {
    this.activeCurrentItem();
  },
};
</script>
