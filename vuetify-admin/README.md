# vuetify-admin

> Clean admin template base on vuetify ui framework

## Build Setup

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

## Guide

### Navigation

`./src/router/nav.json`

```
[
  {
    "name": "<Nav name show in side bar>",
    "icon": "keyboard_arrow_up", // icon default
    "icon-alt": "keyboard_arrow_down", // icon nav items expanded
    "items": [ // children
      {
          "name": "<Nav name show in side bar>",
          "path": "<route path>",
          "component": "<component name in src/pages direction",
          "icon": "<pre icon>",
          "meta": { "requireAuth": true } // meta for more information to beforeRoute
      }
    ]
  },
]
```

### Alert

#### Add to Layout

```
<template>
  <v-app>

    ....

    <v-snackbar
      :color="snackbarColor"
      :value="snackbar"
      @input="closeSnackbar"
    >
      {{ snackbarText }}
      <v-btn dark flat @click.native="closeSnackbar">x</v-btn>
    </v-snackbar>
  </v-app>
</template>

<script>
export default {
  ...

  computed: {
    snackbar() {
      return this.$store.state.g.snackbar;
    },
    snackbarText() {
      return this.$store.state.g.snackbarText;
    },
    snackbarColor() {
      return this.$store.state.g.snackbarColor;
    },
  },

  methods: {
    closeSnackbar() {
      this.$store.dispatch('alert_close');
    },
  },
</script>
```

#### Use
```
this.$store.dispatch('alert_success', msg);
this.$store.dispatch('alert_error', msg);
this.$store.dispatch('alert_warn', msg);
this.$store.dispatch('alert_close');
```

