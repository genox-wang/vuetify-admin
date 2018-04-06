<template>
  <v-container grid-list-md>
    <v-layout row wrap>
      <v-flex xs12>
        <v-card>
          <v-btn color="success" @click="create_dialog = true">Create</v-btn>
        </v-card>
      </v-flex>
    </v-layout>
    <v-layout row wrap>
      <v-flex xs12>
        <v-card>
          <v-card-title>
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
            v-model="tmp"
          >
            <template slot="items" slot-scope="props">
              <td>{{ props.item.ID }}</td>
              <td>{{ props.item.username }}</td>
              <td>{{ props.item.display_name }}</td>
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
    <user-create-dialog v-model="createDialog"/>
    <user-update-dialog v-model="updateDialog" :itemData="selectedItemData"/>
    <delete-line-dialog
      v-model="deleteDialog"
      title="Delete Channel"
      content="sure to delete this line?"
      @delete="onDelete"
    />
  </v-container>

</template>

<script>

import userCreateDialog from '@/dialogs/user_create_dialog';
import userUpdateDialog from '@/dialogs/user_update_dialog';
import deleteLineDialog from '@/dialogs/delete_line_dialog';

export default {
  components: {
    userCreateDialog,
    userUpdateDialog,
    deleteLineDialog,
  },
  data() {
    return {
      createDialog: false,
      updateDialog: false,
      deleteDialog: false,
      selectedItemData: {},
      search: '',
      tmp: [],
      headers: [
        { text: 'ID', value: 'ID', align: 'left' },
        { text: 'Username', value: 'user_name', align: 'left' },
        { text: 'DisplayName', value: 'display_name', align: 'left' },
        { text: 'Action', value: '', align: 'left' },
      ],
    };
  },
  computed: {
    items() {
      return this.$store.state.user.users;
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
      this.$store.dispatch('delete_user', this.selectedItemData.ID)
        .catch(() => {
        });
    },
  },

  mounted() {
    this.$store.dispatch('get_all_users').catch(() => {});
  },
};

</script>
