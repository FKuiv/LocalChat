import type { ChatMessage } from '$lib/types';
import { writable } from 'svelte/store';

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
