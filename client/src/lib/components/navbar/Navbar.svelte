<script lang="ts">
	import Icon from '@iconify/svelte';
	import { page } from '$app/stores';
	import type { iPlayer } from '$lib/types/player';
	$: player = $page.data?.player as iPlayer;

	let showLinkList = true;
	const openLinkList = () => (showLinkList = true);
	const closeLinkList = () => (showLinkList = false);
	const links = [
		{ href: 'auth/register', showWhenLoggedIn: false, name: 'Register' },
		{ href: 'auth/login', showWhenLoggedIn: false, name: 'Login' },
		{ href: '/me', showWhenLoggedIn: true, name: 'Profile' },
		{ href: '/favorites', showWhenLoggedIn: true, name: 'Favorites Agents' }
	];

	$: authedLinks = links.filter((link) => link.showWhenLoggedIn === true);
	$: unauthLinks = links.filter((link) => link.showWhenLoggedIn === false);
</script>

<nav class="flex justify-between items-center p-5">
	<a href="/" class="flex gap-3 text-primary text-3xl font-extrabold font-museo">
		<Icon icon="simple-icons:valorant" />
		<span class="text-white">VALGENTS</span>
	</a>
	<header>
		<!-- svelte-ignore a11y-no-static-element-interactions -->
		{#if player}
			<div
				class="relative z-[10] flex gap-4 items-center bg-primary/70 px-3 py-1"
				on:mouseenter={openLinkList}
				on:mouseleave={closeLinkList}
			>
				<div class="flex items-center gap-3">
					<span class="font-museo font-bold">@{player.username}</span>
					<img
						alt="Your profile"
						class="w-[70px] aspect-square rounded-full object-cover"
						src={player.avatar}
					/>
				</div>

				{#if showLinkList && player}
					<ul class="absolute bottom-0 translate-y-full inset-x-0 bg-neutral-800">
						{#each authedLinks as link}
							<li>
								<a href={link.href}>{link.name}</a>
							</li>
						{/each}
					</ul>
				{/if}
			</div>
		{:else}
			<ul class="flex gap-3 ">
				{#each unauthLinks as link}
					<li class="font-museo hover:underline">
						<a href={link.href}>{link.name}</a>
					</li>
				{/each}
			</ul>
		{/if}
	</header>
</nav>
