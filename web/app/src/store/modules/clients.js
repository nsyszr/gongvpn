import ClientService from "@/services/ClientService";

// initial state
const state = {
  clients: []
};

// getters
const getters = {
  clients(state) {
    return state.clients;
  }
};

// actions
const actions = {
  getClients({ commit }) {
    return new Promise((resolve, reject) => {
      ClientService.getClients()
        .then(r => r.data)
        .then(data => {
          commit("setClients", data);
          resolve(data);
        })
        .catch(err => {
          reject(err);
        });
    });
  },
  createClient({ commit }, client) {
    return new Promise((resolve, reject) => {
      ClientService.createClient(client)
        .then(r => r.data)
        .then(data => {
          commit("addClient", data);
          resolve(data);
        })
        .catch(err => {
          reject(err);
        });
    });
  }
};

// mutations
const mutations = {
  setClients(state, payload) {
    state.clients = payload;
  },
  addClient(state, payload) {
    state.clients.push(payload);
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
