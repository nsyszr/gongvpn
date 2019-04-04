<template>
  <v-container fluid grid-list-xl>
    <v-layout wrap align-center>
      <v-flex xs12>
        <h1>Home</h1>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  computed: {
    ...mapGetters({
      loading: "loading"
    })
  },
  created() {
    // Do nothing
  }
};
</script>

<style>
</style>