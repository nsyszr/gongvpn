<template>
  <v-card class="elevation-2">
    <v-toolbar class="elevation-0" card color="transparent">
      <v-toolbar-title>Clients</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>filter_list</v-icon>
      </v-btn>
      <v-btn icon @click="refresh">
        <v-icon>refresh</v-icon>
      </v-btn>
      <v-menu bottom left>
        <template v-slot:activator="{ on }">
          <v-btn icon v-on="on">
            <v-icon>more_vert</v-icon>
          </v-btn>
        </template>
        <v-list>
          <v-list-tile @click="addApplicationModal=true">
            <v-list-tile-title>Add</v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-menu>
    </v-toolbar>
    <v-container fluid grid-list-xl style="padding: 0;">
      <v-layout row wrap align-center>
        <v-flex xs12>
          <v-data-table
            v-model="selected"
            :headers="headers"
            :items="clients"
            :pagination.sync="pagination"
            :loading="loading"
            select-all
            item-key="id"
            class="elevation-0"
          >
            <template v-slot:headers="props">
              <tr>
                <th>
                  <v-checkbox
                    :input-value="props.all"
                    :indeterminate="props.indeterminate"
                    primary
                    hide-details
                    @click.stop="toggleAll"
                  ></v-checkbox>
                </th>
                <!-- <th style="padding: 0;" class="text-xs-left">&nbsp;</th> -->
                <th
                  v-for="header in props.headers"
                  :key="header.text"
                  :class="['column sortable', pagination.descending ? 'desc' : 'asc', header.value === pagination.sortBy ? 'active' : '']"
                  class="text-xs-left"
                  @click="changeSort(header.value)"
                >
                  {{ header.text }}
                  <v-icon small>arrow_upward</v-icon>
                </th>
              </tr>
            </template>
            <template v-slot:items="props">
              <tr :active="props.selected" @click="props.selected = !props.selected">
                <td>
                  <v-checkbox :input-value="props.selected" primary hide-details></v-checkbox>
                </td>
                <td class="text-xs-left pl-0-pr-0">
                  <v-icon
                    v-if="props.item.connectionStatus == 'Online'"
                    :color="props.item.connectionStatus == 'Online' ? 'green' : 'red'"
                  >{{props.item.connectionStatus == 'Online' ? 'swap_horiz' : 'clear'}}</v-icon>
                </td>
                <td class="text-xs-left" style="width: 20%" nowrap>
                  <a class="link" href="#">{{ props.item.name }}</a>
                </td>
                <td class="text-xs-left" style="width: 20%" nowrap>
                  <v-icon
                    style="padding-right: 8px;"
                  >{{props.item.type == 'Router' ? 'router' : 'computer'}}</v-icon>
                  {{ props.item.type }}
                </td>
                <!-- <td
                  class="text-xs-left"
                  style="width: 16.6%;"
                  nowrap
                  :class="[props.item.connectionStatus === 'Online' ? 'online' : 'offline']"
                >
                  <span>{{ props.item.connectionStatus }}</span>
                </td>-->
                <td
                  class="text-xs-left"
                  style="width: 20%;"
                  nowrap
                >{{ props.item.lastConnectionAt }}</td>
                <td class="text-xs-left" style="width: 20%;" nowrap>{{ props.item.ipv4Address }}</td>
                <td class="text-xs-left" style="width: 20%;" nowrap>{{ props.item.status }}</td>
              </tr>
            </template>
          </v-data-table>
        </v-flex>
      </v-layout>
    </v-container>
  </v-card>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  data: () => ({
    pagination: {
      sortBy: "name"
    },
    selected: [],
    headers: [
      { text: "Conn.", value: "connectionStatus", class: "pl-0-pr-0" },
      {
        text: "Name",
        align: "left",
        value: "name"
      },
      { text: "Type", value: "type" },
      // { text: "Connection Status", value: "connectionStatus" },
      { text: "Connected Since", value: "lastConnectionAt" },
      { text: "VPN IPv4-Address", value: "ipv4Address" },
      { text: "Status", value: "status" }
    ]
  }),
  computed: {
    ...mapGetters({
      loading: "loading",
      clients: "clients/clients"
    })
  },
  methods: {
    toggleAll() {
      if (this.selected.length) this.selected = [];
      else this.selected = this.clients.slice();
    },
    changeSort(column) {
      if (this.pagination.sortBy === column) {
        this.pagination.descending = !this.pagination.descending;
      } else {
        this.pagination.sortBy = column;
        this.pagination.descending = false;
      }
    }
  }
};
</script>

<style scope>
td > .link {
  text-decoration: none;
}
.pl-0-pr-0 {
  padding-left: 0;
  padding-right: 0;
}
</style>
