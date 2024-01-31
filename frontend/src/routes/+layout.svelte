<script lang="ts">
	import '../app.postcss';
	import '../index.css';
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';

	let client = new QueryClient();
</script>

<QueryClientProvider {client}>
	<slot />
</QueryClientProvider>
