<script lang="ts">
	import Icon from '@iconify/svelte';
	import NavLink from './NavLink.svelte';
	import links from './links';
</script>

<nav class="flex justify-between items-center p-5">
	<a href="/" class="flex gap-3 text-primary text-3xl font-extrabold font-museo">
        <Icon icon="simple-icons:valorant" />
		<span class="text-white">VALGENTS</span> 
	</a>
	<header>
		<ul class="flex gap-4">
			{#each links as link}
				<NavLink {link} />
			{/each}
		</ul>
	</header>
</nav>
