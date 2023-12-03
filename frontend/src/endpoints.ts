import axios from "axios";

export const baseUrl = "http://localhost:8000";

export enum UserEndpoints {
  base = `${baseUrl}/user`,
  getAll = `${baseUrl}/users`,
  login = `${baseUrl}/login`,
  logout = `${baseUrl}/logout`,
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
  headers: {
    UserId: localStorage.getItem("UserId"),
    Session: localStorage.getItem("Session"),
  },
});
