<template>
  <v-app id="inspire" :dark="themeIsDark">
    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="primary">
                <v-toolbar-title>Login form</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-form>
                  <v-text-field
                    prepend-icon="person"
                    name="username"
                    label="Username"
                    type="text"
                    v-model="username"
                  >
                  </v-text-field>
                  <v-text-field
                    prepend-icon="lock"
                    name="password"
                    label="Password"
                    type="password"
                    v-model="password"
                  >
                  </v-text-field>

                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                  color="primary"
                  :loading="loading"
                  @click.native="doLogin"
                  :disabled="loading"
                >
                  Login
                  <span slot="loader" class="custom-loader">
                    <v-icon light>cached</v-icon>
                  </span>
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
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
  data() {
    return {
      username: '',
      password: '',
      loading: false,
    };
  },


  computed: {
    token() {
      return this.$store.state.user.token;
    },
    snackbar() {
      return this.$store.state.g.snackbar;
    },
    snackbarText() {
      return this.$store.state.g.snackbarText;
    },
    snackbarColor() {
      return this.$store.state.g.snackbarColor;
    },

    themeIsDark() {
      return this.$store.state.g.themeIsDark;
    },
  },


  methods: {
    doLogin() {
      this.$store.dispatch('login', { username: this.username, password: this.password })
        .then(
          () => {
            this.redirect();
          },
        )
        .catch(() => {
        });
    },

    redirect() {
      if (!this._.isEmpty(this.token)) {
        if (this.$route.query.redirect) {
          this.$router.push(this.$route.query.redirect);
        } else {
          this.$router.push('/');
        }
      }
    },

    closeSnackbar() {
      this.$store.dispatch('alert_close');
    },
  },

  mounted() {
    if (this.$store.state.user.token) {
      this.$router.push('/');
    }
  },
};
</script>

<style>
.custom-loader {
  animation: loader 1s infinite;
  display: flex;
}
@-moz-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@-webkit-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@-o-keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
@keyframes loader {
  from {
    transform: rotate(0);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
