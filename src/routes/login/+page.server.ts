export const actions = {
	login: async ({ cookies, request }) => {
		const data = await request.formData();
		console.log('THE login event', data);
	},
	register: async ({ cookies, request }) => {
		const data = await request.formData();
		console.log('THE register event', data);
	}
};
