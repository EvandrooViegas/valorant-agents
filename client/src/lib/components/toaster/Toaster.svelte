<script lang="ts">
	import { toaster } from "../../../lib/stores/toaster";
	import Item from "./Item.svelte";
    toaster.add({ title: "Hi" })
</script>

<section class="absolute right-4 bottom-4 space-y-3 max-h-[80vh] overflow-auto">
    {#each $toaster as item (item.id)}
    <Item {item} />
{/each}
</section>
