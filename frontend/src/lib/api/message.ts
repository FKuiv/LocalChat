import type { Message, UpdateMessage } from '@/lib/types/message';
import { api } from '.';
import { MessageEndpoints } from './endpoints';
import type { AxiosResponse } from 'axios';

export const createMessage = (messageData: Message): Promise<AxiosResponse<Message>> => {
	return api
		.post<Message>(MessageEndpoints.base(), messageData)
		.then((response) => response)
		.catch((error) => {
			return error;
		});
};

export const getAllMessages = (): Promise<AxiosResponse<Message[]>> => {
	return api
		.get<Message[]>(MessageEndpoints.getAll())
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getMessageById = (id: string): Promise<AxiosResponse<Message>> => {
	return api
		.get<Message>(MessageEndpoints.byId(id))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const updateMessage = (
	id: string,
	messageData: UpdateMessage
): Promise<AxiosResponse<Message>> => {
	return api
		.put<Message>(MessageEndpoints.byId(id), messageData)
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const deleteMessage = (id: string): Promise<AxiosResponse<string>> => {
	return api
		.delete<string>(MessageEndpoints.byId(id))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};

export const getMessagesByGroup = (
	groupId: string | undefined,
	messageAmount: number
): Promise<AxiosResponse<Message[]>> => {
	return api
		.get<Message[]>(MessageEndpoints.getByGroup(groupId, messageAmount))
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};
