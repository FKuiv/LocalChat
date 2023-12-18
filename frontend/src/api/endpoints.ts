export const baseUrl = "http://localhost:8000";

export const UserEndpoints = {
  base: () => `${baseUrl}/user`,
  getAll: () => `${baseUrl}/users`,
  byId: (userId: string) => `${baseUrl}/user/${userId}`,
  login: () => `${baseUrl}/login`,
  logout: () => `${baseUrl}/logout`,
  delete: () => `${baseUrl}/user/delete`,
  profilepic: () => `${baseUrl}/profilepic`,
};

export const GroupEndpoints = {
  base: () => `${baseUrl}/group`,
  getAll: () => `${baseUrl}/groups`,
  getAllUserGroups: () => `${baseUrl}/groups/user`,
  byId: (groupId: string) => `${baseUrl}/groups/${groupId}`,
};

export const MessageEndpoints = {
  base: () => `${baseUrl}/message`,
  getAll: () => `${baseUrl}/messages`,
  byId: (messageId: string) => `${baseUrl}/message/${messageId}`,
};

export enum WebsocketEndpoints {
  base = `${baseUrl}/ws`,
  refresh = `${baseUrl}/ws/refresh`,
}
