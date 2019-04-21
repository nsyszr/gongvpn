<template>
  <v-app>
    <!-- App Loading -->
    <v-layout row justify-center v-if="appLoading">
      <v-dialog v-model="appLoading" persistent fullscreen>
        <v-container fluid fill-height class="loading-dialog">
          <v-layout row justify-center align-center>
            <v-progress-circular indeterminate :size="70" :width="7" color="primary"></v-progress-circular>
          </v-layout>
        </v-container>
      </v-dialog>
    </v-layout>
    <!-- End of App Loading -->

    <!-- Fatal Error Modal -->

    <v-dialog v-model="hasFatalError" persistent>
      <v-card>
        <v-card-title class="headline">A server error occurred!</v-card-title>
        <v-card-text>
          <p class="subheading">Try reloading the app or contact our support.</p>
          <div class="error--text">Error details:</div>
          <div style="padding: 16px; border: 1px solid #ff0000;">
            <samp>{{fatalError}}</samp>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="error" flat @click="logout">Exit</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- End of Fatal Error Modal -->

    <v-navigation-drawer clipped v-model="drawer" fixed app dark>
      <v-list dense>
        <v-list-tile to="/">
          <v-list-tile-action>
            <v-icon>home</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Home</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile to="/users">
          <v-list-tile-action>
            <v-icon>devices</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>VPN-Clients</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar clipped-left color="white" light fixed app class="elevation-2">
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>
        <img class="toolbar-icon" src="./assets/icom_logo.svg" alt>
      </v-toolbar-title>
      <v-toolbar-title class="title">{{title}}</v-toolbar-title>
      <!-- <v-menu offset-y bottom left v-if="currentProject">
        <v-btn slot="activator" small flat style="text-transform: none;">
          <v-icon left style="margin-right:4px; margin-left:-4px;">terrain</v-icon>
          {{currentProject.name}}
          <v-icon right style="margin-left:8px; margin-right:-4px;">arrow_drop_down</v-icon>
        </v-btn>
      </v-menu>-->
      <v-spacer></v-spacer>
      <!-- <v-toolbar-items>
        <v-btn flat>Marktplatz</v-btn>
      </v-toolbar-items>
      <v-menu offset-y bottom left>
        <v-btn slot="activator" icon>
          <v-icon>apps</v-icon>
        </v-btn>
        <v-list two-line subheader>
          <v-subheader inset>Aktive Abonnements</v-subheader>
          <v-list-tile v-for="app in apps" :key="app.title" avatar @click="openUrl(app.url)">
            <v-list-tile-avatar>
              <v-icon :class="[app.iconClass]">{{ app.icon }}</v-icon>
            </v-list-tile-avatar>
            <v-list-tile-content>
              <v-list-tile-title>{{ app.title }}</v-list-tile-title>
              <v-list-tile-sub-title>{{ app.subtitle }}</v-list-tile-sub-title>
            </v-list-tile-content>
            <v-list-tile-action>
              <v-btn icon ripple>
                <v-icon color="grey lighten-1">info</v-icon>
              </v-btn>
            </v-list-tile-action>
          </v-list-tile>
        </v-list>
      </v-menu>
      <v-btn icon>
        <v-icon>notifications</v-icon>
      </v-btn>-->
      <v-menu offset-y bottom left dark>
        <v-btn slot="activator" icon>
          <v-icon>account_circle</v-icon>
        </v-btn>
        <v-list dark>
          <v-list-tile to="/profile">
            <v-list-tile-title>Profile</v-list-tile-title>
          </v-list-tile>
          <v-list-tile @click="logout()" v-if="!isDevMode">
            <v-list-tile-title>Sign out</v-list-tile-title>
          </v-list-tile>
          <v-list-tile @click="true" v-if="isDevMode">
            <v-list-tile-title>Switch User (Dev-Mode)</v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-menu>
    </v-toolbar>
    <v-content>
      <!-- Error -->
      <v-container fluid pt-0 pb-0 style="background-color: #fff;" v-if="hasError">
        <v-layout align-center justify-space-between row wrap fill-height>
          <v-flex>
            <div style="padding-top: 16px; padding-bottom: 16px;">
              <v-icon color="error">warning</v-icon>
              <span
                class="subheading"
                style="padding-left:8px; white-space:normal; word-break: break-all;"
              >{{error}}</span>
            </div>
          </v-flex>
          <v-flex>
            <div class="text-xs-right">
              <v-btn flat color="primary" @click="clearError">Dismiss</v-btn>
            </div>
          </v-flex>
        </v-layout>
      </v-container>
      <!-- End of Error -->

      <!-- <transition>
      <keep-alive>-->
      <router-view></router-view>
      <!-- </keep-alive>
      </transition>-->
    </v-content>
    <!-- <v-footer dark fixed app>
      <span>&copy; 2019 INSYS MICROELECTRONICS GmbH</span>
      <v-spacer></v-spacer>
      <a href="https://www.insys-icom.com/legal-notice" _target="_blank">Legal Notice</a>
      <a href="https://www.insys-icom.com/data-protection" _target="_blank">Data protection</a>
      <a href="https://www.insys-icom.com/products/customer-services" _target="_blank">Hotline</a>
    </v-footer>-->
  </v-app>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

export default {
  name: "App",
  props: {
    source: String
  },
  data: function() {
    return {
      drawer: null,
      apps: [
        {
          icon: "devices",
          iconClass: "grey lighten-1 white--text",
          title: "icom Operations Suite",
          subtitle: "Instanz mustermann (ab8dhsjy9)",
          url: "http://localhost:8083"
        },
        {
          icon: "device_hub",
          iconClass: "grey lighten-1 white--text",
          title: "icom Connectivity Suite - VPN",
          subtitle: "Instanz mustermann (1586)",
          url: "https://connectivity.insys-icom.de"
        }
      ]
    };
  },
  computed: {
    ...mapGetters({
      appLoading: "appLoading",
      hasError: "hasError",
      error: "error",
      hasFatalError: "hasFatalError",
      fatalError: "fatalError",
      isAuthenticated: "isAuthenticated",
      currentProject: "account/currentProject"
    }),
    isDevMode: () => {
      return (
        process.env.VUE_APP_DEV_MODE && process.env.VUE_APP_DEV_MODE === "yes"
      );
    },
    title: () => {
      var title = "Smart IoT VPN";
      if (
        process.env.VUE_APP_DEV_MODE &&
        process.env.VUE_APP_DEV_MODE === "yes"
      ) {
        title = title + " (Dev-Mode)";
      }
      return title;
    }
  },
  methods: {
    ...mapActions({ logout: "logout", clearError: "clearError" }),
    openUrl: function(url) {
      window.open(url, "_blank");
    }
  },
  created() {
    if (this.isAuthenticated) {
      this.$store.dispatch("setAppLoading", false);
    }
  }
};
</script>

<style>
.toolbar-icon {
  margin-top: 8px;
  height: 60px;
}

.nav {
  background-color: #262626;
}

.nav-active {
  background-color: #00a5db;
  color: #fff;
}

.theme--light.application {
  background-color: #e0e0e0;
}

.v-footer {
  padding-left: 16px;
  padding-right: 16px;
}

.v-footer > a {
  color: #ffffff;
  text-decoration: unset;
  padding-left: 32px;
}
.v-footer > a:hover {
  color: #00a5db;
}

.theme--dark.v-footer {
  /*background-color: rgb(38, 38, 38, 0.9);*/
  background-color: #262626;
}

.loading-dialog {
  background-color: #fff;
  /*opacity: 0.85;*/
}
</style>
