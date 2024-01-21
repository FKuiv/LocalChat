import type { WsRefreshMessage } from '../types/message';
import { api } from '.';
import { WebsocketEndpoints } from './endpoints';

export const refreshWebsocket = (message: WsRefreshMessage) => {
	api
		.post(WebsocketEndpoints.refresh, message)
		.then((response) => response)
		.catch((error) => {
			console.error('Error refreshing websocket:', error);
			throw error;
		});
};
