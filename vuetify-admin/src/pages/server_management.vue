<template>
  <v-container grid-list-md>
    <v-layout row wrap>
      <v-flex xs12>
        <v-card>
          <v-btn color="success" @click="createDialog = true">Create</v-btn>
        </v-card>
      </v-flex>
    </v-layout>
    <v-layout row wrap>
      <v-flex xs12>
        <v-card>
          <v-card-title>
            Servers Table
            <v-spacer></v-spacer>
            <v-text-field
              append-icon="search"
              label="Search"
              single-line
              hide-details
              v-model="search"
            ></v-text-field>
          </v-card-title>
          <v-data-table
            :headers="headers"
            :items="items"
            :search="search"
          >
            <template slot="items" slot-scope="props">
              <td>{{ props.item.id }}</td>
              <td>{{ props.item.name }}</td>
              <td>{{ props.item.token }}</td>
              <td width="180">
                <v-btn
                  small
                  fab
                  dark
                  color="cyan"
                  @click.native.stop="doEdit(props.item)"
                >
                  <v-icon>edit</v-icon>
                </v-btn>
                <v-btn
                  small
                  fab
                  dark
                  color="pink"
                  @click.native.stop="doDelete(props.item)"
                >
                  <v-icon>delete</v-icon>
                </v-btn>
              </td>
            </template>
            <v-alert slot="no-results" :value="true" color="error" icon="warning">
              Your search for "{{ search }}" found no results.
            </v-alert>
          </v-data-table>
        </v-card>
      </v-flex>
    </v-layout>
    <server-create-dialog v-model="createDialog"/>
    <server-update-dialog v-model="updateDialog" :itemData="selectedItemData"/>
    <delete-line-dialog
      v-model="deleteDialog"
      title="Delete Server"
      content="sure to delete this line?"
      @delete="onDelete"
    />
  </v-container>

</template>

<script>

import serverCreateDialog from '@/dialogs/server_create_dialog';
import serverUpdateDialog from '@/dialogs/server_update_dialog';
import deleteLineDialog from '@/dialogs/delete_line_dialog';

export default {
  components: {
    serverCreateDialog,
    serverUpdateDialog,
    deleteLineDialog,
  },
  data() {
    return {
      createDialog: false,
      updateDialog: false,
      deleteDialog: false,
      selectedItemData: {},
      search: '',
      headers: [
        { text: 'ID', value: 'id', align: 'left' },
        { text: 'Name', value: 'name', align: 'left' },
        { text: 'Token', value: 'token', align: 'left' },
        { text: 'Action', value: '', align: 'left' },
      ],
    };
  },
  computed: {
    items() {
      return this.$store.state.server.servers;
    },
  },
  methods: {
    doEdit(itemData) {
      this.updateDialog = true;
      this.selectedItemData = itemData;
    },
    doDelete(itemData) {
      this.deleteDialog = true;
      this.selectedItemData = itemData;
    },

    onDelete() {
      this.$store.dispatch('delete_server', this.selectedItemData.id)
        .catch(() => {
        });
    },
  },
  mounted() {
    this.$store.dispatch('get_all_servers').catch(() => {});
  },
};

</script>
