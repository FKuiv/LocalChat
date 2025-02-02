import { source } from 'sveltekit-sse';
import { readChatMessageIdentifier } from '$lib/enums';
// import { addMessage } from '$lib/stores/chatStore';

export const load = async () => {
	const value = source(`/api/chat`).select(readChatMessageIdentifier);
	value.json().subscribe((message) => {
		if (!message) return;
		console.log('new meesasage', message);
		// addMessage(message);
	});
};
