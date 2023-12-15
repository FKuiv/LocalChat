import { GroupRequest, UpdateGroup } from "@/types/group";
import { api } from ".";
import { GroupEndpoints } from "./endpoints";

export const createGroup = (groupData: GroupRequest) => {
  return api
    .post(GroupEndpoints.base(), groupData)
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error creating group:`, error);
      throw error;
    });
};

export const getAllGroups = () => {
  return api
    .get(GroupEndpoints.getAll())
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error getting all groups:`, error);
      throw error;
    });
};

export const getGroupById = (id: string) => {
  return api
    .get(GroupEndpoints.byId(id))
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error getting group by id:`, error);
      throw error;
    });
};

export const updateGroup = (id: string, groupData: UpdateGroup) => {
  return api
    .put(GroupEndpoints.byId(id), groupData)
    .then((response) => response)
    .catch((error) => {
      console.error(`Error updating group:`, error);
      throw error;
    });
};

export const deleteGroup = (id: string) => {
  return api
    .delete(GroupEndpoints.byId(id))
    .then((response) => response)
    .catch((error) => {
      console.error(`Error deleting group:`, error);
      throw error;
    });
};
