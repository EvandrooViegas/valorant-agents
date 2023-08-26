<script lang="ts">
	import { createContext } from '../../utils/createContext';
	import Field from './Field.svelte';
	import Label from './Label.svelte';
	import type { InputContext, InputType } from './types';

	export let value: string | { url: string, path: string };
	export let label: string = '';
	export let name: string;
	export let containerClassName: string = '';
	export let errors: Map<string, string>;
	export let isFileUploaded = false;

	$: error = errors.get(name);
	const setIsFileUploaded = (bool:boolean) => isFileUploaded = bool

	const className =
		'w-full p-3 bg-transparent text-white outline-1 outline-neutral-500 outline-dashed hover:outline-primary focus:outline-primary';
	const type: InputType = $$restProps?.type || 'text';

	const updateValue:InputContext["updateValue"] = (nValue) => {
		if(type !== "file") return value = nValue.value 
		else {
			const ass = nValue as { url: string; value: string }; // Type assertion
			value = { url: ass.url, path: ass.value  }
		}
	}
	const context = createContext<InputContext>('input');
	context.set({
		type,
		label,
		name,
		className,
		setIsFileUploaded,
		updateValue
	});
</script>

<fieldset class={`flex flex-col gap-2 ${containerClassName}`}>
	<Label {isFileUploaded} />
	<Field />
	{#if error}
		<span data-testid={`${name}-error`} class="text-primary text-sm" >{error}</span>
	{/if}
</fieldset>
