<script lang="ts">
	import Button from '$components/button/Button.svelte';
	import type { InputType } from './types';

	export let value: string;
	export let label: string = '';
	export let name: string;
	export let containerClassName: string = '';

	const type: InputType = $$restProps?.type || 'text';
	const clsasName =
		'w-full p-3 bg-transparent text-white outline-1 outline-neutral-500 outline-dashed hover:outline-primary focus:outline-primary';

	const onInputChange = (e: Event) => {
		const inputValue = (e.target as HTMLInputElement).value;
		value = inputValue;
	};

	const onFileChange = (e: Event) => {
		const fileInput = e.target as HTMLInputElement;
		if (fileInput.files?.[0]) {
			value = URL.createObjectURL(fileInput.files[0]);
		}
	};

	const onChange = (e: Event) => {
		switch (type) {
			case 'image': {
				onFileChange(e);
				break;
			}
			default: {
				onInputChange(e);
				break;
			}
		}
	};
	const props = {
		type,
		'on:change': onChange,
		'on:input': onInputChange,
		id: name,
		class: clsasName,
		...$$restProps
	};
</script>

<fieldset class={`flex flex-col gap-2 ${containerClassName}`}>
	{#if label}
		<label for={name} class="text-sm font-semibold">{label}</label>
	{/if}
	{#if label && type === 'file'}
		<label for={name} class="bg-primary w-fit px-4 py-2 font-mono font-semibold">
            Upload a file
        </label>
	{/if}

	{#if type === 'textarea'}
		<textarea {...props} />
	{:else if type === 'file'}
		<input {...props} class="hidden" />
	{:else}
		<input {...props} />
	{/if}
</fieldset>
