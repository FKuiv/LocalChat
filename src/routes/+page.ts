import { source } from 'sveltekit-sse';
import { readChatMessageIdentifier } from '$lib/enums';
import type { PageLoad } from './$types';
// import { addMessage } from '$lib/stores/chatStore';

export const load: PageLoad = async ({ data }) => {
	const value = source(`/api/chat`).select(readChatMessageIdentifier);
	value.subscribe((message) => {
		console.log('READING new message:', message);
		// addMessage(message);
	});
	return data;
};
