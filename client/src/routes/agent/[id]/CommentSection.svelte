<script lang="ts">
	import type { iComment } from "$lib/types/comment";
	import Comment from "./Comment.svelte";
	import CommentForm from "./CommentForm.svelte";
    export let comments:iComment[] = []
</script>
<section>
	<h6 class="text-2xl font-museo">Comments</h6>
    <CommentForm  />
    <div class="flex flex-col gap-3">
        {#each comments as comment}
        <Comment {comment} />
    {/each}
    </div>
</section>