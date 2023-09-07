<script lang="ts">
	import type { HTMLAttributes, HTMLButtonAttributes } from 'svelte/elements';
	import { cva, type VariantProps } from 'class-variance-authority';
	import Icon from '@iconify/svelte';
	const button = cva('button', {
		variants: {
			intent: {
				primary: 'bg-primary',
				border: 'bg-transparent border border-dashed border-neutral-500 font-mono w-fit',
				underline: 'bg-transparent underline hover:text-primary',
				ghost: 'bg-neutral-600/60 font-mono w-fit font-semibold'
			},
			size: {
				small: 'px-2 py-1',
				medium: 'px-4 py-1.5'
			}
		},
		compoundVariants: [
			{ intent: 'primary', size: 'medium', class: 'font-bold font-mono w-fit transition-all' }
		]
	});

	interface $$Props extends HTMLButtonAttributes, VariantProps<typeof button> {
		dataTestId?: string;
		loading?: boolean;
	}
	export let intent: $$Props['intent'] = 'primary';
	export let size: $$Props['size'] = 'medium';
	export let dataTestId: $$Props['dataTestId'] = '';
	$: console.log($$props);
	$: loading = $$props.loading;
</script>

<button
	type="button"
	{...$$props}
	disabled={loading}
	class={button({
		intent,
		size,
		class:
			$$props.class +
			' flex gap-2 items-center disabled:brightness-150 hover:disabled:cursor-not-allowed'
	})}
	data-testid={dataTestId}
>
	{#if loading}
		<Icon icon="line-md:loading-loop" />
	{/if}
	<slot class="w-8" name="icon" />
	<slot />
</button>
