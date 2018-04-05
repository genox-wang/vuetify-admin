<template>
  <base-dialog
    :value = "value"
    @input="updateValue(arguments[0])"
    :title="title"
    cancelText="cancel"
    okText="delete"
    @ok="doDelete"
  >
    {{ content }}
  </base-dialog>
</template>

<script>
import BaseDialog from '@/components/base_dialog';

export default {
  components: {
    BaseDialog,
  },
  props: {
    value: false,
    title: '',
    content: '',
  },
  methods: {
    updateValue(value) {
      this.$emit('input', value);
    },
    doDelete() {
      this.$emit('delete');
      this.updateValue(false);
      // this.$store.dispatch('delete_channel', this.itemData.id)
      //   .catch(() => {
      //   });
      // this.updateValue(false);
    },
  },
};
</script>
