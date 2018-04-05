<template>
  <v-speed-dial
      direction="left"
    >
      <v-btn
        slot="activator"
        color="primary"
        dark
        fab
        hover
        small
      >
        <v-icon>style</v-icon>
      </v-btn>
      <v-btn
        v-for="(color, id) in colors"
        :key="id"
        fab
        dark
        small
        :color="color.css"
        @click.native.stop="changeThemeColor(color)"
      >
      </v-btn>
    </v-speed-dial>
</template>

<script>
import { MUTATION_G_UPDATE_THEME_COLOR } from '@/store/g/mutations_types';

export default {

  data() {
    return {
      colors: [
        {
          field: 'red',
          css: 'red',
        },
        {
          field: 'pink',
          css: 'pink',
        },
        {
          field: 'purple',
          css: 'purple',
        },
        {
          field: 'blue',
          css: 'blue',
        },
        {
          field: 'lightBlue',
          css: 'light-blue',
        },
        {
          field: 'cyan',
          css: 'cyan',
        },
        {
          field: 'teal',
          css: 'teal',
        },
        {
          field: 'green',
          css: 'green',
        },
        {
          field: 'deepOrange',
          css: 'deep-orange',
        },
        {
          field: 'blueGrey',
          css: 'blue-grey',
        },
        {
          field: 'grey',
          css: 'grey',
        },
      ],
    };
  },
  methods: {
    changeThemeColor(color) {
      this.$store.commit(MUTATION_G_UPDATE_THEME_COLOR, color);
    },
  },
};
</script>
