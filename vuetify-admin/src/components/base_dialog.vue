<template>
    <v-dialog
      :value="value"
      @input="updateValue(arguments[0])"
      transition="dialog-right-transition"
      :overlay=false
      @keyup.enter="okClicked"
    >
      <v-card>
        <v-toolbar dark color="primary">
          <v-spacer></v-spacer>
          <v-btn fab small dark color="success" @click.native="okClicked(arguments[0])">
            <v-icon>save</v-icon>
          </v-btn>
        </v-toolbar>
        <v-container>
          <slot ref="slot"></slot>
        </v-container>
      </v-card>
    </v-dialog>
</template>

<script>
export default {
  props: {
    value: false,
    child: null,
  },

  methods: {
    updateValue(value) {
      this.$emit('input', value);
    },
    okClicked(e) {
      this.$emit('ok', e);
    },
    setEnterEscKeyUp() {
      const self = this;
      document.onkeyup = (e) => {
        // 兼容FF和IE和Opera
        const event = e || window.event;
        const key = event.which || event.keyCode || event.charCode;
        if (key === 13) { // enter
          self.okClicked();
        } else if (key === 27) { // esc
          self.updateValue(false);
        }
      };
    },
    removeEnterEscKeyUp() {
      document.onkeyup = null;
    },
  },
  watch: {
    value(cur) {
      if (cur) {
        this.setEnterEscKeyUp();
        const child = this.child;

        this.$nextTick(() => {
          if (child.$refs.focus) {
            child.$refs.focus.focus();
          }
        });
        this.$emit('showed');
      } else {
        this.removeEnterEscKeyUp();
      }
    },
  },
};
</script>

<style lang="stylus">
.dialog-right-transition
  &-enter, &-leave-to
    transform: translateX(100%)
</style>
