<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';
</script>

<div class="flex h-screen">
	<aside class="flex w-1/4 justify-between">
		<div class="grow p-4">here are chats</div>
		<Separator orientation="vertical" />
	</aside>
	<main class="flex grow flex-col-reverse items-center p-4 pb-10">
		<form class="flex w-full justify-between">
			<Input placeholder="Type something" />
			<Button>Send</Button>
		</form>
	</main>
</div>
