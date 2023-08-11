<script lang="ts">
	import type { Option } from './type';

	export let option: Option;
	export let toggleOption: (option: Option) => void;
	export let isSelectedOption: (option: Option) => boolean;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	class={`flex gap-4 items-center ${
		isSelectedOption(option) ? 'bg-neutral-700' : 'bg-transparent'
	} hover:bg-neutral-700 px-4 py-2 cursor-pointer`}
	on:click={() => toggleOption(option)}
>
	<img alt="Role" src={option.image} class="w-[40px]" />
	<span class="font-semibold">{option.name}</span>
</div>
