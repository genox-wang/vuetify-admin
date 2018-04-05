<template>
  <base-dialog
    :value="value"
    @input="updateValue(arguments[0])"
    title="Update Server"
    okText="Save"
    cancelText="Cancel"
    @ok="update"
    @showed="showed"
    >
      <v-layout row wrap>
        <v-flex xs12>
          <v-text-field
            label="Name"
            placeholder="server name"
            v-model="itemData.name"
            @keyup.enter="update"
            ref="focus"
          ></v-text-field>
        </v-flex>
        <v-flex xs12>
          <v-text-field
            label="Token"
            placeholder="token defined by server"
            v-model="itemData.token"
            @keyup.enter="update"
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
    itemData: {},
  },
  methods: {
    updateValue(value) {
      this.$emit('input', value);
    },
    update() {
      this.$store.dispatch('update_server', this.itemData)
        .catch(() => {
        });
      this.updateValue(false);
    },
    showed() {
      this.$nextTick(this.$refs.focus.focus);
    },
  },
};
</script>
