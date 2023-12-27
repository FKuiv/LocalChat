const url = "localhost:8000";
export const baseUrl = `http://${url}`;

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
  byId: (groupId: string | undefined) => `${baseUrl}/group/${groupId}`,
};

export const MessageEndpoints = {
  base: () => `${baseUrl}/message`,
  getAll: () => `${baseUrl}/messages`,
  byId: (messageId: string) => `${baseUrl}/message/${messageId}`,
  getByGroup: (groupId: string | undefined, messageAmount: number) =>
    `${baseUrl}/message/${groupId}/${messageAmount}`,
};

export enum WebsocketEndpoints {
  base = `ws://${url}/ws`,
  refresh = `${baseUrl}/ws/refresh`,
}
