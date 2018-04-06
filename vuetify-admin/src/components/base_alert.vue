<template>
  <v-snackbar
      :color="snackbarColor"
      :value="snackbar"
      @input="closeSnackbar"
      @click.native="closeSnackbar"
    >
      {{ snackbarText }}
    </v-snackbar>
</template>

<script>

export default {
  computed: {
    snackbar() {
      return this.$store.state.g.snackbar;
    },
    snackbarText() {
      return this.$store.state.g.snackbars[0] ? this.$store.state.g.snackbars[0].text : '';
    },
    snackbarColor() {
      return this.$store.state.g.snackbars[0] ? this.$store.state.g.snackbars[0].color : '';
    },
  },

  methods: {
    closeSnackbar() {
      this.$alert.close();
      this.$nextTick(() => {
        this.$alert.next();
      });
    },
  },
};
</script>
