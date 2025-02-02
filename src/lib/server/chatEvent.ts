import { chatEvent } from '$lib/enums';
import { EventEmitter } from 'events';

export const chat = new EventEmitter();

function delay(milliseconds: number) {
	return new Promise(function run(resolve) {
		setTimeout(resolve, milliseconds);
	});
}

export async function startMessaging() {
	for (let i = 0; i < 10; i++) {
		chat.emit(chatEvent, `Hello, world! ${i}`);
		await delay(2000);
	}
}
