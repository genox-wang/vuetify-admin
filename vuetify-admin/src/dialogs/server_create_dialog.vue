<template>
  <base-dialog
    :value="value"
    @input="updateValue(arguments[0])"
    title="Create Server"
    okText="Create"
    cancelText="Cancel"
    @ok="create"
    @showed="showed"
    >
      <v-layout row wrap>
        <v-flex xs12>
          <v-text-field
            label="Name"
            placeholder="server name"
            v-model="name"
            @keyup.enter="create"
            ref="focus"
          ></v-text-field>
        </v-flex>
        <v-flex xs12>
          <v-text-field
            label="Token"
            placeholder="token defined by server"
            v-model="token"
            @keyup.enter="create"
          ></v-text-field>
        </v-flex>
      </v-layout>
  </base-dialog>
</template>

<script>
import baseDialog from '@/components/base_dialog';

export default {
  components: {
    baseDialog,
  },
  props: {
    value: false,
  },
  data() {
    return {
      name: '',
      token: '',
    };
  },
  methods: {
    updateValue(value) {
      this.$emit('input', value);
    },
    create() {
      this.$store.dispatch('create_server', {
        name: this.name,
        token: this.token,
      })
        .catch(() => {
        });
      this.updateValue(false);
    },
    showed() {
      this.$nextTick(this.$refs.focus.focus);
      this.name = '';
    },
  },
};
</script>
