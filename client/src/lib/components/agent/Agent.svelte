<script lang="ts">
	import type { iAgent } from '$lib/types/agents';
	export let agent: iAgent;
</script>

<a
	class="transition-all flex flex-col md:flex-row gap-4 items-center border border-neutral-600 border-dashed hover:border-primary p-5"
	href={`/agent/${agent.uuid}`}
>
	<img alt="Agent Profile" src={agent.displayIcon} class="w-[160px]" />
	<div>
		<span class="text-2xl font-bold">{agent.displayName}</span>
		<p class="text-sm text-neutral-300">{agent.description}</p>
	</div>
</a>
