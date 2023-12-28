import { Message } from "./message";
import { User } from "./user";

export type GroupRequest = {
  name: string;
  user_ids: string[];
  admins: string[];
  isDm: boolean;
};

export type UpdateGroup = {
  name: string;
  user_ids: string[];
  admins: string[];
};

export type Usernames = { [key: string]: string };

export type Group = {
  id: string;
  name: string;
  usernames: Usernames;
  created_at: Date;
  updated_at: Date;
  users: User[];
  messages: Message[];
  admins: string[];
  is_dm: boolean;
};

export const defaultGroup: Group = {
  id: "",
  name: "",
  usernames: {},
  created_at: new Date(),
  updated_at: new Date(),
  users: [],
  messages: [],
  admins: [],
  is_dm: false,
};
