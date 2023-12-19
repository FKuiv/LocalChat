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

export type Group = {
  id: string;
  name: string;
  created_at: Date;
  updated_at: Date;
  users: User[];
  messages: Message[];
  admins: string[];
  isDm: boolean;
};

export const defaultGroup: Group = {
  id: "",
  name: "",
  created_at: new Date(),
  updated_at: new Date(),
  users: [],
  messages: [],
  admins: [],
  isDm: false,
};
