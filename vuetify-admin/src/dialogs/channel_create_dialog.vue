<template>
  <base-dialog
    :value="value"
    @input="updateValue(arguments[0])"
    title="Create Channel"
    okText="Create"
    cancelText="Cancel"
    @ok="create"
    @showed="showed"
    >
      <v-layout row wrap>
        <v-flex xs12>
          <v-text-field
            label="Name"
            placeholder="channel Name"
            v-model="name"
            @keyup.enter="create"
            ref="focus"
          ></v-text-field>
        </v-flex>
        <v-flex xs12>
            <v-expansion-panel>
            <v-expansion-panel-content  key="servers">
              <div slot="header"><strong>Servers Setting</strong></div>
              <v-card>
                <v-card-text>
                  <v-container fluid>
                    <v-layout row wrap>
                      <v-flex v-for="server in servers" :key="server.id" xs12 sm4 md4>
                        <v-checkbox
                          :label="server.name"
                          v-model="server.enabled"
                          color="success"
                          hide-details
                        ></v-checkbox>
                      </v-flex>
                    </v-layout>
                  </v-container>
                </v-card-text>
              </v-card>
            </v-expansion-panel-content>
          </v-expansion-panel>
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
      servers: [],
    };
  },
  methods: {
    updateValue(value) {
      this.$emit('input', value);
    },
    create() {
      const newServers = [];
      this.servers.forEach((e) => {
        if (e.enabled) {
          newServers.push({ id: e.id });
        }
      });
      this.$store.dispatch('create_channel', {
        name: this.name,
        servers: newServers,
      })
        .catch(() => {
        });
      this.updateValue(false);
    },
    showed() {
      this.$nextTick(this.$refs.focus.focus);
      this.name = '';
      const self = this;
      this.$store.dispatch('get_all_servers')
        .then(() => {
          self.servers = self.$store.state.server.servers;
        })
        .catch(() => {});
    },
  },
};
</script>
