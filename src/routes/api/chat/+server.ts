import { readChatMessageIdentifier, writeChatMessageIdentifier } from '$lib/enums';
import { messageEventEmitter } from '$lib/server/chatEvent';
import type { RequestHandler } from '@sveltejs/kit';
import { produce } from 'sveltekit-sse';

export const POST: RequestHandler = () => {
	return produce(async function start({ emit }) {
		const send = (content: string) => {
			const { error } = emit(readChatMessageIdentifier, content);
			if (error) {
				return cancel();
			}
		};

		const cancel = () => {
			messageEventEmitter.removeListener(writeChatMessageIdentifier, send);
		};

		messageEventEmitter.on(writeChatMessageIdentifier, send);
		return cancel;
	});
};
