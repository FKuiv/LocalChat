import { writeChatMessageIdentifier } from '$lib/enums.js';
import { messageEventEmitter } from '$lib/server/chatEvent.js';
import { error, fail } from '@sveltejs/kit';

export const actions = {
	default: async (event) => {
		console.log('formeevetn', event);
		const data = await event.request.formData();
		const message = data.get('message');

		if (!message) {
			return error(400, 'Invalid request');
		}

		try {
			messageEventEmitter.emit(writeChatMessageIdentifier, message);
			return {
				status: 200,
				body: {
					success: true
				}
			};
		} catch (err) {
			console.error('Error creating message:', err);
			return fail(500, { error: 'Failed to create message' });
		}
	}
};
