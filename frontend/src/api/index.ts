import axios from "axios";
import { baseUrl } from "./endpoints";
import * as usersApi from "./user";
import * as groupsApi from "./group";
import * as messagesApi from "./message";
import * as websocketApi from "./websocket";

export const api = axios.create({
  baseURL: baseUrl,
  withCredentials: true,
});

export const ping = () => {
  return api
    .get("/")
    .then((response) => response)
    .catch((error) => {
      console.error(`Error pinging server:`, error);
      throw error;
    });
};

export { usersApi, groupsApi, messagesApi, websocketApi };
