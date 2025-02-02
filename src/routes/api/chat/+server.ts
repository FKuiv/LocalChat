import { chatEvent } from '$lib/enums';
import { chat } from '$lib/server/chatEvent';
import type { RequestHandler } from '@sveltejs/kit';
import { produce } from 'sveltekit-sse';

export const POST: RequestHandler = () => {
	return produce(async function start({ emit }) {
		const send = () => {
			const { error } = emit(chatEvent, 'Hello, world!');
			if (error) {
				return cancel();
			}
			console.log('Ran send');
		};

		const cancel = () => {
			chat.removeListener(chatEvent, send);
		};

		chat.on(chatEvent, send);
		return cancel;
	});
};
