<script lang="ts">
	import type { SuperValidated } from 'sveltekit-superforms';
	import { superForm } from 'sveltekit-superforms/client';
	import type { LoginSchema } from '$lib/types/schemas';
	import TextInput from '$lib/components/formInputs/TextInput.svelte';
	import PasswordInput from '$lib/components/formInputs/PasswordInput.svelte';

	export let data: SuperValidated<LoginSchema>;

	const { form, errors, enhance, constraints } = superForm(data);
</script>

<form use:enhance class="flex flex-col gap-y-8" method="POST">
	<TextInput
		name="username"
		label="Username"
		placeholder="John Mary"
		bind:value={$form.username}
		errors={$errors.username}
		constraints={$constraints.username}
	/>
	<PasswordInput
		name="password"
		label="Password"
		placeholder="Password"
		bind:value={$form.password}
		errors={$errors.password}
		constraints={$constraints.password}
	/>
	<div class="flex justify-end items-center">
		<button class="btn variant-glass-primary" type="submit">Submit</button>
	</div>
</form>
