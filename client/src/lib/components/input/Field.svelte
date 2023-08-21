<script lang="ts">
	import storage from '$lib/services/storage';
	import { toaster } from '$lib/stores/toaster';
	import { createContext } from '$lib/utils/createContext';
	import type { InputContext } from './types';

	const context = createContext<InputContext>('input').get();
	let { type, updateValue, name, className, setIsFileUploaded } = context;
	let lastImageUploaded: string | undefined;
	const onInput = (e: Event) => {
		const value = (e.target as HTMLInputElement).value;
		updateValue({ type, value }, name);
	};

	const onFileChange = async (e: Event) => {
		const fileInput = e.target as HTMLInputElement;
		const file = fileInput.files?.[0];
		if (!file) return;

		try {
			const image = await storage.uploadImage(file);
			if (!image) return;
			toaster.add({ title: 'Image Upload Successfully', description: 'Beautify Avatar Btw' });

			const imagePath = URL.createObjectURL(file)
			updateValue({ type: "file", value: imagePath, url: image.url }, name);

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
		type,
		id: name,
		class: className
	};
</script>
<button type="button" on:click={handleRemoveImage}>
	Remove 
</button>
{#if type === 'textarea'}
	<textarea {...props} on:input={onInput} />
{:else if type === 'file'}
	<input {...props} on:change={onFileChange} class="hidden" />
{:else}
	<input {...props} on:input={onInput} />
{/if}
