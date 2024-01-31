import { type UserMap, type Login, type User } from '@/lib/types/user';
import { api } from '.';
import { UserEndpoints } from './endpoints';
import type { AxiosResponse } from 'axios';

export const getAllUsers = (): Promise<AxiosResponse<User[]>> => {
	return api
		.get<User[]>(UserEndpoints.getAll())
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getAllUsersMap = (): Promise<AxiosResponse<UserMap>> => {
	return api
		.get<UserMap>(UserEndpoints.getAllMap())
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getUserById = (userId: string): Promise<AxiosResponse<User>> => {
	return api
		.get<User>(UserEndpoints.byId(userId))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const createUser = (loginInfo: Login): Promise<AxiosResponse<User>> => {
	return api
		.post<User>(UserEndpoints.base(), loginInfo)
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const deleteUser = (): Promise<AxiosResponse<string>> => {
	return api
		.delete<string>(UserEndpoints.delete())
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const updateUser = (updatedData: Login): Promise<AxiosResponse<User>> => {
	return api
		.put<User>(UserEndpoints.base(), updatedData)
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const loginUser = (loginData: Login): Promise<AxiosResponse<string>> => {
	return api
		.post<string>(UserEndpoints.login(), loginData)
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const logoutUser = (): Promise<AxiosResponse<string>> => {
	return api
		.get<string>(UserEndpoints.logout())
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getUsername = (userId: string): Promise<AxiosResponse<string>> => {
	return api
		.get<string>(UserEndpoints.username(userId))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const uploadUserPicture = (formData: FormData): Promise<AxiosResponse<string>> => {
	return api
		.post<string>(UserEndpoints.picture(), formData, {
			headers: { 'Content-Type': 'multipart/form-data' }
		})
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getUserPicture = (userId: string): Promise<AxiosResponse<string>> => {
	return api
		.get<string>(UserEndpoints.getPicture(userId))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};
