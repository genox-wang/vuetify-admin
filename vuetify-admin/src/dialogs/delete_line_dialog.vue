<template>
  <base-dialog
    :value = "value"
    @input="updateValue(arguments[0])"
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
    content: '',
  },
  methods: {
    updateValue(value) {
      this.$emit('input', value);
    },
    doDelete() {
      this.$emit('delete');
      this.updateValue(false);
    },
  },
};
</script>
