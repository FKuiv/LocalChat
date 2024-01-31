import type { Group, GroupRequest, UpdateGroup } from '@/lib/types/group';
import { api } from '.';
import { GroupEndpoints } from './endpoints';
import type { AxiosResponse } from 'axios';

export const createGroup = (groupData: GroupRequest): Promise<AxiosResponse<Group>> => {
	return api
		.post<Group>(GroupEndpoints.base(), groupData)
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getAllGroups = (): Promise<AxiosResponse<Group[]>> => {
	return api
		.get<Group[]>(GroupEndpoints.getAll())
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getGroupById = (id: string | undefined): Promise<AxiosResponse<Group>> => {
	return api
		.get<Group>(GroupEndpoints.byId(id))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const updateGroup = (id: string, groupData: UpdateGroup): Promise<AxiosResponse<Group>> => {
	return api
		.put<Group>(GroupEndpoints.byId(id), groupData)
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const deleteGroup = (id: string): Promise<AxiosResponse<string>> => {
	return api
		.delete<string>(GroupEndpoints.byId(id))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getAllUserGroups = (): Promise<AxiosResponse<Group[]>> => {
	return api
		.get<Group[]>(GroupEndpoints.getAllUserGroups())
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const uploadGroupPicture = (
	groupId: string,
	formData: FormData
): Promise<AxiosResponse<string>> => {
	return api
		.post<string>(GroupEndpoints.picture(groupId), formData, {
			headers: { 'Content-Type': 'multipart/form-data' }
		})
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getGroupPicture = (groupId: string): Promise<AxiosResponse<string>> => {
	return api
		.get<string>(GroupEndpoints.picture(groupId))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const groupExistingByUserIdsAndAdmins = (
	userIds: string[],
	admins: string[]
): Promise<AxiosResponse<Group[]>> => {
	return api
		.post<Group[]>(GroupEndpoints.existing(), { user_ids: userIds, admins: admins })
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};
