import { writable } from 'svelte/store';

export type ChatMessage = {
	id: string;
	content: string;
	timestamp: Date;
};

export const messages = writable<ChatMessage[]>([]);

export function addMessage(content: string) {
	messages.update((msgs) => [
		...msgs,
		{
			id: crypto.randomUUID(),
			content,
			timestamp: new Date()
		}
	]);
}
