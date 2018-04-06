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

#### Use
```
// in components
this.$alert.success(msg)
this.$alert.error(msg)
this.$alert.warn(msg)
this.$alert.info(msg)
this.$alert.close(msg)

// global
Vue.$alert.success(msg)
Vue.$alert.error(msg)
Vue.$alert.warn(msg)
Vue.$alert.info(msg)
Vue.$alert.close(msg)
```

