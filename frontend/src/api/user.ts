import { Login } from "@/types/user";
import { api } from ".";
import { UserEndpoints } from "./endpoints";

export const getAllUsers = () => {
  return api
    .get(UserEndpoints.getAll())
    .then((response) => response.data)
    .catch((error) => {
      console.error("Error fetching all users:", error);
      throw error;
    });
};

export const getUserById = (userId: string) => {
  return api
    .get(UserEndpoints.byId(userId))
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error fetching user with ID ${userId}:`, error);
      throw error;
    });
};

export const createUser = (loginInfo: Login) => {
  return api
    .post(UserEndpoints.base(), loginInfo)
    .then((response) => response)
    .catch((error) => {
      console.error("Error creating user:", error);
      throw error;
    });
};

export const deleteUser = () => {
  return api
    .delete(UserEndpoints.delete())
    .then((response) => response)
    .catch((error) => {
      console.error(`Error deleting user:`, error);
      throw error;
    });
};

export const updateUser = (updatedData: Login) => {
  return api
    .put(UserEndpoints.base(), updatedData)
    .then((response) => response)
    .catch((error) => {
      console.error(`Error updating user:`, error);
      throw error;
    });
};

export const loginUser = (loginData: Login) => {
  return api
    .post(UserEndpoints.login(), loginData)
    .then((response) => response)
    .catch((error) => {
      console.error(`Error logging in user:`, error);
      throw error;
    });
};

export const logoutUser = () => {
  return api
    .get(UserEndpoints.logout())
    .then((response) => response)
    .catch((error) => {
      console.log("Error logging out user:", error);
      throw error;
    });
};
