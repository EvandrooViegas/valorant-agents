<script lang="ts">
	import Button from '../../../lib/components/button/Button.svelte';
	import Input from '../../../lib/components/input/Input.svelte';
	import { createContext } from '../../../lib/utils/createContext';
	import z from 'zod';
	import player from '../../../lib/services/player';
	import { toaster } from '$lib/stores/toaster';
	import type { LoginFormContext } from './types';

	let username: string;
	let password: string;
	let errors = new Map<string, string>();
	const formSchema = z.object({
		username: z
			.string()
			.min(5, 'Username must have at least 5 digits')
			.max(50, 'Username must not have more than 50 digits'),
		password: z
			.string()
			.min(5, 'Password must have at least 5 digits')
			.max(50, 'Password must not have more than 50 digits')
	});
	const onSubmit = async (e: SubmitEvent) => {
		try {
			e?.preventDefault();
			const form = { username, password };

			const response = await player.login(form);
		} catch (error) {
			toaster.add({
				title: 'Error',
				description: 'Could not create an account, try again later',
				type: 'error'
			});
		}
	};

	const contextKey = 'login-form';
	const context = createContext<LoginFormContext>(contextKey);
	const contextValue = { formSchema };
	context.set(contextValue);
</script>

<form
	on:submit={onSubmit}
	enctype="multipart/form-data"
	class="mx-auto max-w-[800px] flex flex-col md:grid md:grid-cols-2 gap-4"
	data-testid="form"
>
	<Input {errors} formContext={context} label="Username" bind:value={username} name="username" />
	<Input
		{errors}
        showGeneratePasswordBtn={false}
		formContext={context}
		label="Password"
		bind:value={password}
		type="password"
		name="password"
	/>
	<Button type="submit" class="col-span-2" id="submit-btn">Submit</Button>
</form>
