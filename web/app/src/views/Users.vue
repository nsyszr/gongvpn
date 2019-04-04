<template>
  <v-container fluid grid-list-xl>
    <v-layout wrap="true" align-center>
      <v-flex xs12 d-flex>
        <users-card></users-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import UsersCard from "@/components/UsersCard.vue";

export default {
  components: {
    UsersCard
  },
  created() {
    this.$store.dispatch("clients/getClients");
  }
};
</script>

<style>
</style>
