<script lang="ts">
	import type Nil from '../../../lib/types/nil.js';
	import type { iAgent } from '../../../lib/types/agents.js';

	export let data;
	const agent: Nil | iAgent = data?.agent;
	const backgroundClass = `background-image: url(${agent.background})`
</script>

{#if agent}

<div class="space-y-8">
	<section>
		<header
			class={`p-6 relative flex md:flex-row flex-col items-center justify-between bg-fill  bg-center`}
			style={backgroundClass}
		>
			<div class="absolute inset-0 gradient-overlay z-[0]" />
			<div class="relative">
				<img src={agent.displayIcon} class="w-[250px]" alt="Agent Disaply Icon" />
				<h4 class="text-6xl font-black font-museo bg-primary/80 px-4">{agent.displayName}</h4>
			</div>
			<div class="relative flex flex-col items-center gap-4">
				<img src={agent.role.displayIcon} class="w-[90px]" alt="Role Display Icon" />
				<span class="text-3xl font-semibold">{agent.role.displayName}</span>
			</div>
		</header>
	</section>

	<section>
		<h6 class="text-2xl font-museo">Description</h6>
		<p>{agent.description}</p>
	</section>

	<section>
		<h6 class="text-2xl font-museo">Abilities</h6>
		<div class="flex">
			<div class="grid md:grid-cols-2 gap-4 flex-wrap">
				{#each agent.abilities as abilitie}
					<div
						class="w-full flex flex-col md:flex-row items-center gap-8 border border-neutral-600 p-5"
					>
						<img src={abilitie.displayIcon} class="w-[80px] object-contain" alt="Display Icon" />
						<div class="flex flex-col gap-2 md:text-left text-center">
							<span class="texl-xl font-bold">
								{abilitie.displayName}
							</span>
							<span class="text-sm">
								{abilitie.description}
							</span>
						</div>
					</div>
				{/each}
			</div>
		</div>
	</section>
</div>	

{:else}
	No agent found
{/if}