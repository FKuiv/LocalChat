import type { Group } from './group';
import type { Message } from './message';

export type Login = {
	username: string;
	password: string;
};

export const defaultLogin: Login = {
	username: '',
	password: ''
};

export type Session = {
	id: string;
	user_id: string;
	created_at: Date;
	updated_at: Date;
	expires_at: Date;
};

export type User = {
	id: string;
	username: string;
	password: string;
	created_at: Date;
	updated_at: Date;
	messages: Message[];
	groups: Group[];
	session: Session;
};
