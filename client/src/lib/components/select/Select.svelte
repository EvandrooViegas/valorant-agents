<script lang="ts">
	import type { Option } from './type';
	//props
	export let options: Option[] = [];
	export let label: string = '';
	export let onSelectChange: (options:Option[]) => undefined | void

	//consts
	let OPTIONS = options;

	//states
	let isDialogOpen = false;

	const isSelectedOption = (option: Option) => {
		return options?.some((opt) => opt.id === option.id);
	};
	const toggleIsDialogOpen = () => {
		isDialogOpen = !isDialogOpen;
	};

	const toggleOption = (option: Option) => {
		const exists = isSelectedOption(option);
		if (exists) {
			options = options.filter((opt) => option.id !== opt.id);
		} else {
			options = [...options, option];
		}
		OPTIONS = [...OPTIONS];
		onSelectChange(options)
	};
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="flex flex-col justify-center gap-2 h-full">
	<span>{label}</span>
	<div
		on:click={toggleIsDialogOpen}
		class="relative px-4 py-2 h-full border border-dashed border-neutral-500"
	>
		{#if options.length > 0}
			<div class="relative overflow-x-hidden space-x-3">
				{#each options as option}
				<span>{option.name}</span>
			{/each}
			<div class="text-center absolute right-0 px-2 inset-y-0 bg-neutral-800">
				{options.length} {" "} Selected
			</div> 
			</div>
		{:else}
			<span>Nothing Selected</span>
		{/if}
		{#if isDialogOpen}
			<div
				on:click={(e) => e.stopPropagation()}
				class="absolute inset-x-0 top-full bg-background py-4 border border-dashed"
			>
				<div class="flex flex-col gap-4">
					{#each OPTIONS as option (option.id)}
						<div
							class={`flex gap-4 items-center ${
								isSelectedOption(option) ? 'bg-neutral-700' : 'bg-transparent'
							} hover:bg-neutral-700 px-4 py-2 cursor-pointer`}
							on:click={() => toggleOption(option)}
						>
							<img alt="Role" src={option.image} class="w-[40px]" />
							<span class="font-semibold">{option.name}</span>
						</div>
					{/each}
				</div>
			</div>
		{/if}
	</div>
</div>
