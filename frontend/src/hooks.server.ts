import { redirect } from '@sveltejs/kit';

const public_paths = ['/login', '/register'];

function isPublicPath(path: string): boolean {
	return public_paths.includes(path);
}

export async function handle({ event, resolve }) {
	let isAuthenticated = false;
	if (event.cookies.get('Session') && event.cookies.get('UserId')) {
		isAuthenticated = true;
	}

	if (!isAuthenticated && !isPublicPath(event.url.pathname)) {
		throw redirect(302, '/login');
	}

	if (isAuthenticated) {
		event.locals.userId = event.cookies.get('UserId');
		event.locals.session = event.cookies.get('Session');
		if (event.url.pathname === '/login') {
			throw redirect(302, '/');
		}
	}

	const response = await resolve(event);
	return response;
}
