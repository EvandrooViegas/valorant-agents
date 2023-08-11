<script lang="ts">
	import Agent from '$components/agent/Agent.svelte';
	import Select from '$components/select/Select.svelte';
	import type { Option } from '$components/select/type.js';
	import { converRolesToOptions, convertOptionsToRoles, getRolesFromAgents } from '$components/select/utils';
	import valorant from '$lib/services/valorant';
	import type { iRole } from '$lib/types/agents';
	import isObjectFalse from '$lib/utils/isObjectFalse';
	import type { Filter } from './types';

	//data
	export let data;
	const agents = data.agents;

	//states
	let filtredAgents = agents;
	let selectedOptions:Option<iRole>[] | null = null
	//handlers
	const onSelectChange = (options: Option[]) => {
		if (!selectedOptions) return;
		selectedOptions = options;
	};

	const filterAgents = async () => {
		const response = await valorant.filterAgents(filters);
		if (response) filtredAgents = [...response];
		else filtredAgents = []
	};
	//
	const roles = getRolesFromAgents(agents)
	const selectOptions = converRolesToOptions(roles)

	//computed
	$: filters = { name: "", roles: convertOptionsToRoles(selectedOptions) } satisfies Filter;
	$: isFiltering = isObjectFalse(filters)

	//state handlers
	const clearFilters = () => {
		filters = { name: '', roles: [] };
		filterAgents()
	};
</script>

{#if agents}{#if filtredAgents}
		<div class="space-y-12">
			<form class="md:grid grid-cols-2 gap-4" on:submit={filterAgents}>
				<fieldset class="flex flex-col gap-2">
					<label for="name" class="font-bold">Name</label>
					<input
						id="name"
						bind:value={filters.name}
						class="px-4 py-2 bg-transparent border border-dashed border-neutral-500 outline-none"
					/>
				</fieldset>
				<fieldset>
					<Select label="Roles" options={selectOptions} {onSelectChange} />
				</fieldset>
				{#if isFiltering}
					<div class="space-x-2">
						<button class=" px-4 py-1.5 bg-primary md:w-fit w-full font-bold font-mono">
							Filter
						</button>
						<button class="underline hover:text-primary" on:click={clearFilters}>
							Remove Filters
						</button>
					</div>
				{/if}
			</form>
			<ul class="grid md:grid-cols-2 grid-cols-1 gap-4">
				<div class="col-span-2">
					<h3 class="text-3xl font-semibold font-museo">Agents:</h3>
				</div>
				{#each filtredAgents as agent}
					<li><Agent {agent} /></li>
				{/each}
			</ul>
		</div>
	{:else}
		<span>No Agents Found</span>
	{/if}
{:else}
	<div>Error Fetching agents</div>
{/if}
