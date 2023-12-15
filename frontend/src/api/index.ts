import axios from "axios";
import { baseUrl } from "./endpoints";

export const api = axios.create({
  baseURL: baseUrl,
  withCredentials: true,
});
