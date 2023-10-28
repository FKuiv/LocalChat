import axios from "axios";

const baseUrl = "http://localhost:8000";

export enum UserEndpoints {
  base = `${baseUrl}/user`,
  getAll = `${baseUrl}/users`,
}

export enum GroupEndpoints {
  base = `${baseUrl}/group`,
  getAll = `${baseUrl}/groups`,
}

export enum MessageEndpoints {
  base = `${baseUrl}/message`,
  getAll = `${baseUrl}/messages`,
}

export const api = axios.create({
  baseURL: baseUrl,
});
