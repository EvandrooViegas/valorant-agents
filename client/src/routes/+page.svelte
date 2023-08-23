<script lang="ts">
	import Agent from '../lib/components/agent/Agent.svelte';
	import valorant from '../lib/services/valorant';
	import FilterForm from './FilterForm.svelte';
	import type { Filter } from './types.js';

	//data
	export let data;
	const agents = data.agents;

	//states
	let filtredAgents = agents;
	//handlers
	const filterAgents = async (filters: Filter) => {
		const response = await valorant.filterAgents(filters);
		if (response) filtredAgents = [...response];
		else filtredAgents = [];
	};

	
</script>

{#if agents}
	{#if filtredAgents && filterAgents.length > 0}
		<div class="space-y-12">
			<FilterForm {agents} {filtredAgents}  onSubmit={filterAgents} />
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
