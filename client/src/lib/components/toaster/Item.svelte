<script lang="ts">
	import type { Item } from "$lib/stores/toaster";
    export let item:Item
</script>

<div class="space-y-1 border border-dashed border-neutral-600 px-5 py-2.5">
    <span class="font-semibold">{item.title}</span>
    {#if item.description}
    <p class="text-sm">{item.description}</p>
    {/if}
</div>  