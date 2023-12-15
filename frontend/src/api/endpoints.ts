export const baseUrl = "http://localhost:8000";

export enum UserEndpoints {
  base = `${baseUrl}/user`,
  getAll = `${baseUrl}/users`,
  login = `${baseUrl}/login`,
  logout = `${baseUrl}/logout`,
  delete = `${baseUrl}/user/delete`,
  profilepic = `${baseUrl}/profilepic`,
}

export enum GroupEndpoints {
  base = `${baseUrl}/group`,
  getAll = `${baseUrl}/groups`,
}

export enum MessageEndpoints {
  base = `${baseUrl}/message`,
  getAll = `${baseUrl}/messages`,
}
