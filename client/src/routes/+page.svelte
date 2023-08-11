<script lang="ts">
	import Agent from '$components/agent/Agent.svelte';
	import Select from '$components/select/Select.svelte';
	import type { Option } from '$components/select/type.js';
	import valorant from '$lib/services/valorant';
	import type { iRole } from '$lib/types/agents';
	import type { Filter } from './types';
	//data
	export let data;

	const agents = data.agents;

	//states
	let filtredAgents = agents;
	let name = '';
	//handlers
	const onSelectChange = (options: Option[]) => {
		if (!selectedOptions) return;
		selectedOptions = options;
	};

	const filterAgents = async () => {
		await valorant.filterAgents(filters)			
	}
	//
	const roles = agents
		?.reduce((curr, agent) => {
			const exists = curr.map((c) => c.displayName).includes(agent.role.displayName);
			return exists ? [...curr] : [...curr, agent.role];
		}, [] as iRole[])
		.filter((r) => r.displayName !== '' || r.displayIcon !== '');

	let selectedOptions = roles?.map((r, idx) => ({
		id: idx,
		name: r.displayName,
		value: r,
		image: r.displayIcon
	})) as Option[];



	//computed
	$: filters = {
		name,
		roles: selectedOptions.map(o => o.value) as iRole[]
	} satisfies Filter

	$: isFiltering = (() => {
		if (!filters) return false;

		return Object.keys(filters).some((k) => {
			const value = filters[k as keyof typeof filters];
			if (typeof value === 'object') {
				return Object.keys(value).length > 0;
			} else return Boolean(value);
		});
	})();
</script>

{#if agents}{#if filtredAgents}
		<div class="space-y-12">
			<form class="md:grid grid-cols-2 gap-4" on:submit={filterAgents}>
				<fieldset class="flex flex-col gap-2">
					<label for="name" class="font-bold">Name</label>
					<input
						id="name"
						bind:value={name}
						class="px-4 py-2 bg-transparent border border-dashed border-neutral-500 outline-none"
					/>
				</fieldset>
				<fieldset>
					<Select label="Roles" options={selectedOptions} {onSelectChange} />
				</fieldset>
				{#if isFiltering}
					<button class=" px-4 py-1.5 bg-primary md:w-fit w-full font-bold font-mono">
						Filter
					</button>
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
