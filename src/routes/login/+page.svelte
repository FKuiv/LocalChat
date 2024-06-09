<script lang="ts">
	import LoginForm from '$lib/components/LoginForm.svelte';
	import { TabGroup, Tab } from '@skeletonlabs/skeleton';
	let tabSet: number = 0;
</script>

<main class="flex flex-col justify-center items-center h-full">
	<form method="POST" action="?/login" class="card p-5 flex flex-col gap-y-8 w-1/3">
		<TabGroup>
			<Tab bind:group={tabSet} name="login" value={0}>Login</Tab>
			<Tab bind:group={tabSet} name="register" value={1}>Register</Tab>
			<svelte:fragment slot="panel">
				{#if tabSet === 0}
					<LoginForm formaction="?/login" formName="login" />
				{:else}
					<LoginForm formaction="?/register" formName="register" />
				{/if}
			</svelte:fragment>
		</TabGroup>
	</form>
</main>
