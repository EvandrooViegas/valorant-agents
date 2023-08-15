<script lang="ts">
	import Button from '$components/button/Button.svelte';
	import Select from '$components/select/Select.svelte';
	import type { Option } from '$components/select/type';
	import {
		converRolesToOptions,
		convertOptionsToRoles,
		getRolesFromAgents
	} from '$components/select/utils';
	import type { iAgent, iRole } from '$lib/types/agents';
	import isObjectFalse from '$lib/utils/isObjectFalse';
	import type { Filter } from './types';

	export let onSubmit: (filters: Filter) => Promise<void>;
	export let name: string = '';
	export let filtredAgents: iAgent[] | null;
	export let agents: iAgent[] | null;

	const roles = getRolesFromAgents(agents);

	let options = converRolesToOptions(roles);
	let selectedOptions: Option<iRole>[] | null = null;

	//computed
	$: selectedRoles = convertOptionsToRoles(selectedOptions);
	$: isFiltering = isObjectFalse({ name, selectedRoles });

	//state handlers
	const clearFilters = () => {
		name = '';
		selectedRoles = [];
		selectedOptions = null;
		filtredAgents = agents;

		// to trigger a re-render on the select component
		options = [...options];
	};
	const onSelectChange = (options: Option[]) => {
		selectedOptions = options;
	};
	const handleSubmit = () => {
		onSubmit({ name, roles: selectedRoles });
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
		<div >
			<Button>Filter </Button>
			<Button intent="underline" on:click={clearFilters}> Remove Filters </Button>
		</div>
	{/if}
</form>
