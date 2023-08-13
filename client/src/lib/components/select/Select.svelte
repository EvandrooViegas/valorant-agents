<script lang="ts">
	import SelectOption from './SelectOption.svelte';
	import type { Option } from './type';
	
	//props
	export let options: Option[];
	export let label: string = '';
	export let onSelectChange: (options: Option[]) => undefined | void;
	export let selectedOptions:  Option[] | null

	//states
	let isDialogOpen = false;

	//handlers
	const opendDialog = () => {
		isDialogOpen = true
	};
	const closeDialog = () => {
		isDialogOpen = false
	}

	const toggleOption = (option: Option) => {
		const exists = isSelectedOption(option);
		if (exists) {
			selectedOptions && (selectedOptions = selectedOptions.filter((opt) => option.id !== opt.id));
		} else {
			selectedOptions
				? (selectedOptions = [...selectedOptions, option])
				: (selectedOptions = [option]);
		}
		options = [...options];
		onSelectChange(selectedOptions || []);
	};
	
	const isSelectedOption = (option: Option) => {
		return Boolean(selectedOptions?.some((opt) => opt.id === option.id))
	};

</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="flex flex-col justify-center gap-2 h-full">
	<span>{label}</span>
	<div
		on:click={opendDialog}
		on:mouseenter={opendDialog}
		on:mouseleave={closeDialog}
		class="relative px-4 py-2 h-full border border-dashed border-neutral-500"
	>
		{#if selectedOptions && selectedOptions?.length > 0}
			<div class="relative overflow-x-hidden space-x-3">
				{#each selectedOptions as option}
					<span>{option.name}</span>
				{/each}
				<div class="text-center absolute right-0 px-2 inset-y-0 bg-neutral-800">
					{selectedOptions.length}
					{' '} Selected
				</div>
			</div>
		{:else}
			<span>Nothing Selected</span>
		{/if}
		{#if isDialogOpen && options}
			<div
				on:mouseleave={closeDialog}
				on:click={(e) => e.stopPropagation()}
				class="absolute inset-x-0 top-full bg-background py-4 border border-dashed"
			>
				<div class="flex flex-col gap-4">
					{#each options as option (option.id)}
						<SelectOption {toggleOption} {isSelectedOption} {option} />
					{/each}
				</div>
			</div>
		{/if}
	</div>
</div>
