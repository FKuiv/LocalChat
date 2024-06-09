<script lang="ts">
	export let formaction = '';
	export let formName = '';
</script>

<div class="flex flex-col gap-y-5">
	<label class="label">
		<span>Username</span>
		<input class="input" type="text" name="username" placeholder="John Mary" />
	</label>
	<label class="label">
		<span>Password</span>
		<input class="input" type="password" name="password" />
	</label>

	<button
		{formaction}
		type="submit"
		class="btn bg-gradient-to-br variant-gradient-secondary-tertiary w-1/4 self-end capitalize"
	>
		{formName}
	</button>
</div>
