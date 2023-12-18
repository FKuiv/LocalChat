export type MessageRequest = {
  group_id: string;
  content: string;
};

export type UpdateMessage = {
  content: string;
};

export type WsRefreshMessage = {
  new_group_id: string;
  clients_to_update: string[];
};

export type Message = {
  id: string;
  user_id: string;
  group_id: string;
  content: string;
  created_at: Date;
  updated_at: Date;
  DeletedAt: boolean | Date;
};
