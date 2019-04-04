import Api from "@/services/Api";

export default {
  getClients() {
    return Api(true).get("/clients");
  },
  createClient(client) {
    return Api(true).post("/clients", client);
  }
};
