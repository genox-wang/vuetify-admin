<template>
   <v-btn icon @click.native.stop="toggleIsDark">
      <v-icon>{{themeIsDark?'brightness_7':'brightness_4'}}</v-icon>
    </v-btn>
</template>

<script>
import { MUTATION_G_UPDATE_THEME_IS_DARK } from '@/store/g/mutations_types';

export default {
  computed: {
    themeIsDark() {
      return this.$store.state.g.themeIsDark;
    },
  },
  methods: {
    toggleIsDark() {
      this.$store.commit(MUTATION_G_UPDATE_THEME_IS_DARK, !this.themeIsDark);
    },
  },
};
</script>
