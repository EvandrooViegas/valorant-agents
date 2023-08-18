<script lang="ts">
	import storage from '$lib/services/storage';
	import { toaster } from '$lib/stores/toaster';
	import { createContext } from '$lib/utils/createContext';
	import type { InputContext } from './types';

	const context = createContext<InputContext>('input').get();
	const { type, updateValue, name, className } = context;

	const onInput = (e: Event) => {
		const value = (e.target as HTMLInputElement).value;
		updateValue(value, name);
	};

	const onFileChange = async (e: Event) => {
		const fileInput = e.target as HTMLInputElement;
		const file = fileInput.files?.[0];
		if (!file) return;
		const form = new FormData();
		form.append('avatar', file);
		try {
			const responseImage = await storage.upload(form)
			if(responseImage) {
				const url = responseImage.url
				toaster.add({ title: "Image Upload Successfully", description: "Beautify Avatar Btw" })
				toaster.add({ title: "Image Upload Successfully", description: "Beautify Avatar Btw" })
				toaster.add({ title: "Image Upload Successfully", description: "Beautify Avatar Btw" })
				updateValue(url, name);
			} 
		} catch (error) {
			console.log('error uploading the image', error);
		}
	};

	const props = {
		type,
		id: name,
		class: className
	};
</script>

{#if type === 'textarea'}
	<textarea {...props} on:input={onInput} />
{:else if type === 'file'}
	<input {...props} on:change={onFileChange} class="hidden" />
{:else}
	<input {...props} on:input={onInput} />
{/if}
