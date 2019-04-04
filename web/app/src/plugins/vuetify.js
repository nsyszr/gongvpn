import Vue from "vue";
import Vuetify from "vuetify/lib";
// import 'vuetify/src/stylus/app.styl'
import "../stylus/main.styl";

const isDevMode =
  process.env.VUE_APP_DEV_MODE && process.env.VUE_APP_DEV_MODE === "yes";

Vue.use(Vuetify, {
  iconfont: "md",
  theme: {
    primary: isDevMode ? "#00A5DB" : "#00A5DB" // FF5722
  }
});
