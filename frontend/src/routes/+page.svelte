<script>
	import { getAllUsers } from '@/api/user';
	import { useQuery } from '@sveltestack/svelte-query';

	let query = useQuery('test', getAllUsers);
	console.log($query);
</script>

<div class="container h-full mx-auto flex justify-center items-center">
	Localchat
	{#if $query.isLoading}
		Loading...
	{:else if $query.isError}
		{$query.error}
	{:else}
		{#each $query.data as user}
			<div>{user.username}</div>
		{/each}
	{/if}
</div>
