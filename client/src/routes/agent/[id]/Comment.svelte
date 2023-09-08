<script lang="ts">
	import { page } from '$app/stores';
	import type { iComment } from '$lib/types/comment';
	import Icon from '@iconify/svelte';
	import type { iPlayer } from '../../../lib/types/player';
	import CommentForm from './CommentForm.svelte';
	import Button from '$lib/components/button/Button.svelte';
	export let comment: iComment;
	let showReplies = false;
	$: player = $page.data?.player as iPlayer 
</script>

<li class="list-none p-4 border border-dashed border-white/10 ">
	<div class="flex flex-col gap-3">
		<header class="flex justify-between">
            <div class="flex items-center gap-2">
                <img src={player.avatar} alt="Player Avatar" class="w-12 aspect-square rounded-full object-cover" />
                <span class="text-sm font-black font-museo">@{player.username}</span>
            </div>
            <span class="text-sm font-black font-museo text-neutral-500">
                Posted on: Dec 22, 2023
            </span>
        </header>
		<p class="">
			{comment.text}
		</p>
	</div>
	<div>
		<div class="pl-4">
            <CommentForm parentComment={comment} />
			{#if comment.replies?.length > 0}
				<button class="flex items-center gap-1 py-2" on:click={() => (showReplies = !showReplies)}>
					<span>
						<Icon
							icon="iconamoon:arrow-right-2"
							class={`transition-all ${showReplies ? 'rotate-90' : 'rotate-0'}`}
						/>
					</span>
					<span class={`transition-all text-xs font-extrabold  ${showReplies ? 'text-primary' : 'text-neutral-500'}`}>
						{comment.replies.length} Replies
					</span>
				</button>
			{/if}
			{#if showReplies}
				<ul class="flex flex-col gap-2">
					{#each comment.replies as comment}
						<svelte:self {comment} />
					{/each}
				</ul>
			{/if}
		</div>
	</div>
	{#if comment.author.id === player.id}
	<div>
		<Button intent="underline">
			Delete Comment
		</Button>
	</div>
	{/if}
</li>
