<script lang="ts">
	import { TabGroup, TabAnchor, AppBar } from '@skeletonlabs/skeleton';
	import BasicForm from './BasicForm.svelte';
	import { page } from '$app/stores';
	import Logo from './Logo.svelte';
	import type { SuperValidated } from 'sveltekit-superforms';
	import type { LoginSchema } from '../types/schemas';

	let tabSet: number = 0;
	export let formRoute = 'login';
	export let data: SuperValidated<LoginSchema>;

	if (formRoute === 'register') {
		tabSet = 1;
	}
</script>

<div class="w-1/2 flex gap-y-5 flex-col">
	<Logo />
	<TabGroup justify="justify-center">
		<TabAnchor
			bind:group={tabSet}
			name="login"
			value={0}
			href="/login"
			selected={$page.url.pathname === '/login'}>Login</TabAnchor
		>
		<TabAnchor
			bind:group={tabSet}
			name="register"
			value={1}
			href="/register"
			selected={$page.url.pathname === '/register'}>Register</TabAnchor
		>
		<svelte:fragment slot="panel">
			<BasicForm {data} />
		</svelte:fragment>
	</TabGroup>
</div>
