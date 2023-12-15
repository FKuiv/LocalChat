import { MessageRequest, UpdateMessage } from "@/types/message";
import { api } from ".";
import { MessageEndpoints } from "./endpoints";

export const createMessage = (messageData: MessageRequest) => {
  return api
    .post(MessageEndpoints.base(), messageData)
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error creating message:`, error);
      throw error;
    });
};

export const getAllMessages = () => {
  return api
    .get(MessageEndpoints.getAll())
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error getting all messages:`, error);
      throw error;
    });
};

export const getMessageById = (id: string) => {
  return api
    .get(MessageEndpoints.byId(id))
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error getting message by id:`, error);
      throw error;
    });
};

export const updateMessage = (id: string, messageData: UpdateMessage) => {
  return api
    .put(MessageEndpoints.byId(id), messageData)
    .then((response) => response)
    .catch((error) => {
      console.error(`Error updating message:`, error);
      throw error;
    });
};

export const deleteMessage = (id: string) => {
  return api
    .delete(MessageEndpoints.byId(id))
    .then((response) => response)
    .catch((error) => {
      console.error(`Error deleting message:`, error);
      throw error;
    });
};
