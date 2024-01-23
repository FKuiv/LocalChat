<script lang="ts">
	import { getAllUsers } from '@/api/user';
	import { useQuery } from '@sveltestack/svelte-query';
	import type { User } from '@/types/user';
	import type { AxiosResponse } from 'axios';

	let query = useQuery<AxiosResponse<User[]>>('test', getAllUsers);
	console.log($query);
</script>

<div class="container h-full mx-auto flex justify-center items-center">
	Localchat
	{#if $query.isLoading}
		Loading...
	{:else if $query.isError}
		{$query.error}
	{:else}
		{#each $query.data?.data ?? [] as user (user.id)}
			<div>{user.username}</div>
		{/each}
	{/if}
</div>
