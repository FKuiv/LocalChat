import { startMessaging } from '$lib/server/chatEvent';
import type { Load } from '@sveltejs/kit';

export const load: Load = () => {
	startMessaging();
};
