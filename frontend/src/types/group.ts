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
  users: boolean | User[];
  messages: boolean | Message[];
  admins: string[];
  isDm: boolean;
};
