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
            Users Table
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
              <td>{{ props.item.username }}</td>
              <td>{{ props.item.display_name }}</td>
              <td width="180">
                <v-btn small fab dark color="cyan"><v-icon>edit</v-icon></v-btn>
                <v-btn small fab dark color="pink"><v-icon>delete</v-icon></v-btn>
              </td>
            </template>
            <v-alert slot="no-results" :value="true" color="error" icon="warning">
              Your search for "{{ search }}" found no results.
            </v-alert>
          </v-data-table>
        </v-card>
      </v-flex>
    </v-layout>
    <user-create-dialog v-model="create_dialog"/>
  </v-container>

</template>

<script>

import userCreateDialog from '@/dialogs/user_create_dialog';

export default {
  components: {
    userCreateDialog,
  },
  data() {
    return {
      create_dialog: false,
      search: '',
      headers: [
        { text: 'ID', value: 'id', align: 'left' },
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
  mounted() {
    this.$store.dispatch('get_all_users').catch(() => {});
  },
};

</script>
