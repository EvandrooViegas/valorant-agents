<script lang="ts">
	import Button from '../../../lib/components/button/Button.svelte';
	import Input from '../../../lib/components/input/Input.svelte';
	import PlayerCard from '../../../lib/components/player-card/PlayerCard.svelte';
	import { createContext } from '../../../lib/utils/createContext';
	import z from 'zod';
	import type { RegisterFormContext } from './types';
	import player from '../../../lib/services/player';
	import validateSchema from '../../../lib/utils/validateSchema';
	import { toaster } from '$lib/stores/toaster';

	let avatar: { path: string; url: string } = { path: '', url: '' };
	let username: string;
	let password: string;
	let description: string;
	let errors = new Map<string, string>();
	let loading = false
	const formSchema = z.object({
		username: z
			.string()
			.min(5, 'Username must have at least 5 digits')
			.max(50, 'Username must not have more than 50 digits'),
		description: z
			.string()
			.min(5, 'Description must have at least 5 digits')
			.max(150, 'Description must not have more than 150 digits'),
		password: z
			.string()
			.min(5, 'Password must have at least 5 digits')
			.max(50, 'Password must not have more than 50 digits'),
		avatar: z.string()
	});
	const onSubmit = async (e: SubmitEvent) => {
		loading = true
		try {
			e?.preventDefault();
			errors = new Map();
			const form = { avatar: avatar.url, username, password, description };
			const result = validateSchema(formSchema, form);
			if (result.type === 'error') {
				errors = result.errors;
				return;
			}
			errors = new Map();

			const response = await player.create(form);
			toaster.add({
				title: 'Error',
				description: 'Could not create an account, try again later',
				type: 'success'
			});
			const id = response?.id;
			if (id) {
				location.replace("/")
			} else {
				toaster.add({
					title: 'Error',
					description: 'An error occurred',
					type: 'error'
				});
			}
		} catch (error) {
			toaster.add({
				title: 'Error',
				description: 'Could not create an account, try again later',
				type: 'error'
			});
		}
		loading = false
	};

	const contextKey = 'register-form';
	const context = createContext<RegisterFormContext>(contextKey);
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
		formContext={context}
		label="Password"
		bind:value={password}
		type="password"
		name="password"
	/>
	<Input
		formContext={context}
		label="Description"
		bind:value={description}
		name="description"
		maxlength="150"
		type="textarea"
		containerClassName="col-span-2"
		{errors}
	/>
	<Input
		formContext={context}
		label="Avatar"
		{errors}
		bind:value={avatar}
		name="avatar"
		type="file"
		accept="image/png, image/jpeg"
	/>
	<Button {loading} type="submit" class="col-span-2" id="submit-btn">Submit</Button>

	<div class="mt-8 col-span-2">
		<span class="font-museo font-semibold text-xl">Player Card Preview: </span>
		<PlayerCard {username} {description} avatar={avatar.path} />
	</div>
</form>
