<script lang="ts">
	import Button from "../../components/button/Button.svelte";
	import { createContext } from "../../utils/createContext";
	import type { InputContext } from "./types";

	export let isFileUploaded: boolean
	const context = createContext<InputContext>('input').get();

	let { label, name, type } = context;

</script>

{#if label}
	<label for={name} class="text-sm font-semibold">{label}</label>
{/if}
{#if label && type === 'file'}
<Button intent={isFileUploaded ? "border" : "primary"} >
	<label for={name}>
			{isFileUploaded ? "Replace Avatar" : "Upload Avatar"}
		</label>
	</Button>
{/if}
