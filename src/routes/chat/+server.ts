import { chat } from '$lib/stores/chat';
import { produce } from 'sveltekit-sse';

export function POST() {
	return produce(async function start({ emit }) {
		const send = () => {
			const { error } = emit('message', 'Hello, world!');
			if (error) {
				return cancel();
			}
		};

		const cancel = () => {
			chat.removeListener('message', send);
		};

		chat.on('message', send);
		return cancel;
	});
}
