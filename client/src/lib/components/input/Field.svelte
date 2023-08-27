<script lang="ts">
	import player from '../../../lib/services/player';
	import storage from '../../services/storage';
	import { toaster } from '../../stores/toaster';
	import { createContext } from '../../utils/createContext';
	import Hide from '../icons/Hide.svelte';
	import Show from '../icons/Show.svelte';
	import type { InputContext } from './types';
	
	const context = createContext<InputContext>('input').get();
	const { type, updateValue, name, className, setIsFileUploaded } = context;
	let lastImageUploaded: string | undefined;
	let fieldValue: string;
	let showPassword: boolean = false;

	const toggleShowPassword = () => {
		showPassword = !showPassword;
	};
	const onInput = (e: Event) => {
		updateValue({ type, value: fieldValue }, name);
	};
	const generatePassword = async () => {
		const password = await player.generatePassword();
		fieldValue = password;
		updateValue({ type: 'password', value: password }, name);
	};
	const onFileChange = async (e: Event) => {
		const fileInput = e.target as HTMLInputElement;
		const file = fileInput.files?.[0];
		if (!file) return;

		try {
			const image = await storage.uploadImage(file);
			if (!image) return;
			toaster.add({ title: 'Image Upload Successfully', description: 'Beautify Avatar Btw' });

			updateValue({ type: 'file', value: URL.createObjectURL(file), url: image.url }, name);

			setIsFileUploaded(true);
			handleRemoveImage();
			lastImageUploaded = image.name;
		} catch (error) {
			console.log('error uploading the image', error);
		}
	};
	const handleRemoveImage = () => {
		if (lastImageUploaded) storage.deleteImage(lastImageUploaded);
	};
	const props = {
		'data-testid': name,
		type,
		id: name,
		class: className
	};
</script>

{#if type === 'textarea'}
	<textarea {...props} bind:value={fieldValue} on:input={onInput} />
{:else if type === 'file'}
	<input {...props} on:change={onFileChange} class="hidden" />
{:else if type === 'password'}
	<div class="relative group">
		{#if showPassword}
			<input {...props} type="text" bind:value={fieldValue} on:input={onInput} />
		{:else}
			<input {...props} bind:value={fieldValue} on:input={onInput} />
		{/if}
		<div
			class="flex gap-1 transition-all opacity-0 group-hover:opacity-100 absolute right-1 top-1/2 -translate-y-1/2"
		>
			<button type="button" class="p-2 bg-neutral-700" on:click={toggleShowPassword}>
				{#if showPassword}
				<Show />
				{:else}
				<Hide />
				{/if}
			</button>
			<button type="button" on:click={generatePassword} class="p-2 bg-neutral-700"
				>Generate Password
			</button>
		</div>
	</div>
{:else}
	<input {...props} bind:value={fieldValue} on:input={onInput} />
{/if}
