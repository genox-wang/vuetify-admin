<template>
   <base-dialog
    :value="value"
    @input="updateValue(arguments[0])"
    title="Update Channel"
    okText="Save"
    cancelText="Cancel"
    @ok="update"
    @showed="showed"
    >
      <v-layout row wrap>
        <v-flex xs12>
          <v-text-field
            label="Name"
            placeholder="channel name"
            v-model="itemData.name"
            @keyup.enter="update"
            ref="focus"
          ></v-text-field>
        </v-flex>
        <v-flex xs12>
            <v-expansion-panel>
            <v-expansion-panel-content key="servers">
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
    itemData: {},
  },
  data() {
    return {
      servers: [],
    };
  },
  methods: {
    updateValue(value) {
      this.$emit('input', value);
    },
    update() {
      const newServers = [];
      this.servers.forEach((e) => {
        if (e.enabled) {
          newServers.push({ id: e.id });
        }
      });
      this.itemData.servers = newServers;
      this.$store.dispatch('update_channel', this.itemData)
        .catch(() => {
        });
      this.updateValue(false);
    },
    showed() {
      this.$nextTick(this.$refs.focus.focus);
      const self = this;
      this.$store.dispatch('get_all_servers')
        .then(() => {
          const servers = self.$store.state.server.servers;
          self.servers = servers.map((e) => {
            let i = 0;
            e.enabled = false;
            if (self.itemData.servers) {
              for (;i < self.itemData.servers.length; i += 1) {
                if (self.itemData.servers[i].id === e.id) {
                  e.enabled = true;
                  break;
                }
              }
            }
            return { id: e.id, name: e.name, enabled: e.enabled };
          });
        })
        .catch(() => {});
    },
  },
};
</script>
