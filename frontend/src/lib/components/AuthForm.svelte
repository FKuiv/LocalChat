<script lang="ts">
	import { TabGroup, Tab, TabAnchor } from '@skeletonlabs/skeleton';
	import BasicForm from './BasicForm.svelte';

	let tabSet: number = 0;
</script>

<TabGroup justify="justify-center">
	<Tab bind:group={tabSet} name="tab1" value={0}>Login</Tab>
	<Tab bind:group={tabSet} name="tab2" value={1}>Create account</Tab>
	<svelte:fragment slot="panel">
		{#if tabSet === 0}
			<BasicForm />
		{:else if tabSet === 1}
			<BasicForm />
		{/if}
	</svelte:fragment>
</TabGroup>
