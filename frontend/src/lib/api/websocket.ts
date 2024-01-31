import type { WsRefreshMessage } from '../types/message';
import { api } from '.';
import { WebsocketEndpoints } from './endpoints';
import type { AxiosResponse } from 'axios';

export const refreshWebsocket = (message: WsRefreshMessage): Promise<AxiosResponse<string>> => {
	api
		.post<string>(WebsocketEndpoints.refresh, message)
		.then((response) => response)
		.catch((error) => {
			throw error;
		});
};
