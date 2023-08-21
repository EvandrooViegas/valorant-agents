<script lang="ts">
	import Button from '$components/button/Button.svelte';
	import Input from '$components/input/Input.svelte';
	import PlayerCard from '$components/player-card/PlayerCard.svelte';
	import { createContext } from '$lib/utils/createContext';
	import z from 'zod';
	import type { RegisterFormContext } from './types';
	import player from '$lib/services/player';

	let avatar: string;
	let username: string;
	let password: string;
	let description: string;
	let errors = new Map<string, string>();

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
		errors.clear();
		e?.preventDefault();
		const form = { avatar, username, password, description };
		const result = formSchema.safeParse(form);
		const hasErrors = 'error' in result;
		if (hasErrors) {
			const fields = Object.keys(form);
			const formErrors = result.error.format();
			fields.forEach((field) => {
				const error = formErrors[field as keyof typeof form]?._errors[0];
				if (error) {
					errors = errors.set(field, error);
				}
			});
		} else {
			const result = await player.create(form);
		}
	};

	const contextKey = 'register-form';
	const context = createContext<RegisterFormContext>(contextKey);
	const contextValue = { formSchema };
	context.set(contextValue);
</script>

<form on:submit={onSubmit} enctype="multipart/form-data" class="grid grid-cols-2 gap-4">
	<Input {errors} formContext={context} label="Username" bind:value={username} name="username" />
	<Input {errors} formContext={context} label="Password" bind:value={password} name="password" />
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
	<Button type="submit" class="col-span-2">Submit</Button>

	<div>
		<span>Preview: </span>
		<PlayerCard {username} {description} {avatar} />
	</div>
</form>
