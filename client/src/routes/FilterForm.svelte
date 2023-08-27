<script lang="ts">
	import Button from '../lib/components/button/Button.svelte';
	import Select from '../lib/components/select/Select.svelte';
	import type { Option } from '../lib/components/select/type';
	import {
		converRolesToOptions,
		convertOptionsToRoles,
		getRolesFromAgents
	} from '../lib/components/select/utils';
	import type { iAgent, iRole } from '../lib/types/agents';
	import isObjectFalse from '../lib/utils/isObjectFalse';
	import type { Filter } from './types';

	export let onSubmit: (filters: Filter) => Promise<void>;
	export let name: string = '';
	export let agents: iAgent[] | null;
	export let resetFilteredAgents: () => void
	const roles = getRolesFromAgents(agents);

	let options = converRolesToOptions(roles);
	let selectedOptions: Option<iRole>[] | null = null;

	//computed
	$: isFiltering = isObjectFalse({ name, selectedRoles: convertOptionsToRoles(selectedOptions) });
	//state handlers
	const clearFilters = () => {
		name = '';
		selectedOptions = null;
		resetFilteredAgents()
		
		// to trigger a re-render on the select component
		options = [...options];
	};
	const onSelectChange = (options: Option[]) => {
		selectedOptions = options;
	};
	const handleSubmit = () => {
		onSubmit({ name, roles: convertOptionsToRoles(selectedOptions) });
	};
</script>

<form class="md:grid grid-cols-2 gap-4" on:submit={handleSubmit}>
	<fieldset class="flex flex-col gap-2">
		<label for="name" class="font-bold">Name</label>
		<input
			id="name"
			bind:value={name}
			class="px-4 py-2 bg-transparent border border-dashed border-neutral-500 outline-none"
		/>
	</fieldset>
	<fieldset>
		<Select label="Roles" {selectedOptions} {options} {onSelectChange} />
	</fieldset>
	{#if isFiltering}
		<div>
			<Button type="submit">Filter</Button>
			<button on:click={clearFilters}>Remove Filters</button>
		</div>
	{/if}
</form>
