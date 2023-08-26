<script lang="ts">
	import storage from '../../services/storage';
	import { toaster } from '../../stores/toaster';
	import { createContext } from '../../utils/createContext';
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

			updateValue({ type: "file", value: URL.createObjectURL(file), url: image.url }, name);
		
			setIsFileUploaded(true);
			handleRemoveImage();
			lastImageUploaded = image.name;
		} catch (error) {
			console.log('error uploading the image', error);
		}
	};
	// unit -> vitest
	// integration
	// e2e
	const handleRemoveImage = () => {
		if (lastImageUploaded) storage.deleteImage(lastImageUploaded);
	};
	const props = {
		"data-testid": name,
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
