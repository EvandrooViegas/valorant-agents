<script lang="ts">
	import type { Link } from "./links";
    export let link: Link
    
</script>

<li>
    <a class="text-semibod hover:underline" href={link.url}>{link.name}</a>
</li>