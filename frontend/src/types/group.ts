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
