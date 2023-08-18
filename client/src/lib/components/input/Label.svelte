<script lang="ts">
	import { createContext } from "$lib/utils/createContext";
	import type { InputContext } from "./types";

	const context = createContext<InputContext>('input').get();
	const { label, name, type } = context;

</script>

{#if label}
	<label for={name} class="text-sm font-semibold">{label}</label>
{/if}
{#if label && type === 'file'}
	<label for={name} class="bg-primary w-fit px-4 py-2 font-mono font-semibold">
		Upload a file
	</label>
{/if}
