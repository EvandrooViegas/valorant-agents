<script lang="ts">
	import { createContext, type Context } from '$lib/utils/createContext';
	import Field from './Field.svelte';
	import Label from './Label.svelte';
	import type { InputContext, InputType } from './types';

	export let value: string;
	export let label: string = '';
	export let name: string;
	export let containerClassName: string = '';
	export let errors:Map<string, string>
	$: error = errors.get(name) 
	

	const className =
	'w-full p-3 bg-transparent text-white outline-1 outline-neutral-500 outline-dashed hover:outline-primary focus:outline-primary';
	const type: InputType = $$restProps?.type || 'text';

	const updateValue = (nValue: any) => {
		value = nValue;
	};
	const context = createContext<InputContext>('input');
	context.set({
		type,
		label,
		name,
		className,
		updateValue
	});

</script>

<fieldset class={`flex flex-col gap-2 ${containerClassName}`}>
	<Label />

	<Field />
	{#if error}
		<span class="text-primary text-sm">{error}</span>
	{/if}
</fieldset>
