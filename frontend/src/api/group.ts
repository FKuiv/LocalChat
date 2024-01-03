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

export const getGroupById = (id: string | undefined) => {
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

export const getAllUserGroups = () => {
  return api
    .get(GroupEndpoints.getAllUserGroups())
    .then((response) => response.data)
    .catch((error) => {
      console.error("Error getting user groups:", error);
      throw error;
    });
};

export const uploadGroupPicture = (groupId: string, formData: FormData) => {
  return api
    .post(GroupEndpoints.picture(groupId), formData, {
      headers: { "Content-Type": "multipart/form-data" },
    })
    .then((response) => response)
    .catch((error) => {
      console.error("Error uploading group picture:", error);
      throw error;
    });
};

export const getGroupPicture = (groupId: string) => {
  return api
    .get(GroupEndpoints.picture(groupId))
    .then((response) => response.data)
    .catch((error) => {
      console.error("Error getting group picture:", error);
      throw error;
    });
};
