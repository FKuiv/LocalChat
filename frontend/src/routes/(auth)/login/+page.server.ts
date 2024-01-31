import { superValidate } from 'sveltekit-superforms/server';
import { LoginSchema } from '@/lib/types/schemas';
import { error } from '@sveltejs/kit';

export const load = async () => {
	const form = await superValidate(LoginSchema);

	return { form };
};

export const actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, LoginSchema);
		console.log('POST', form);

		if (!form.valid) {
			return error(400, 'Invalid form');
		}

		// TODO: Log the user in

		return { form };
	}
};
