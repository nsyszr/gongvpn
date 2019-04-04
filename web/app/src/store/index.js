import Vue from "vue";
import Vuex from "vuex";
import app from "./modules/app";
import session from "./modules/session";
import clients from "./modules/clients";
import createLogger from "@/utils/logger";

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== "production";

export default new Vuex.Store({
  modules: {
    app,
    session,
    clients
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
});
