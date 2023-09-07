<script lang="ts">
	import Icon from '@iconify/svelte';
	import { page } from '$app/stores';
	import type { iPlayer } from '$lib/types/player';
	import playerService from '$lib/services/player';
	$: player = $page.data?.player as iPlayer;

	let showLinkList = false;
	const openLinkList = () => (showLinkList = true);
	const closeLinkList = () => (showLinkList = false);
	$: links = [
		{ href: '/auth/register', showWhenLoggedIn: false, name: 'Register' },
		{ href: '/auth/login', showWhenLoggedIn: false, name: 'Login' },
		{ href: `/player/${player?.username}`, showWhenLoggedIn: true, name: 'Profile' },
		{ href: '/favorites', showWhenLoggedIn: true, name: 'Favorites Agents' }
	];

	$: authedLinks = links.filter((link) => link.showWhenLoggedIn === true);
	$: unauthLinks = links.filter((link) => link.showWhenLoggedIn === false);
	const logout = async () => {
		await playerService.logout()
        document.location.reload()

	}
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
					<ul class="absolute bottom-0 translate-y-full inset-x-0 flex flex-col bg-neutral-800">
						{#each authedLinks as link}
							<li class=" font-museo border-b border-dashed border-transparent hover:border-neutral-600">
								<a  href={link.href} class="p-3 block">{link.name}</a>
							</li>
						{/each}
						<li class="">
							<button on:click={logout} class="font-museo w-full p-4 font-museo border-t border-dashed border-neutral-600">
								Logout
							</button>
						</li>
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