<script lang="ts">
	import PlayerCard from '$components/player-card/PlayerCard.svelte';

	let image: string;
	let username: string;
	let password: string;
	let description: string;

	const onImageLoad = (e: any) => {
		if (e?.target?.files?.[0]) {
			image = URL.createObjectURL(e.target.files[0]);
		}
	};
</script>

<form method="POST" action="?/create" enctype="multipart/form-data" class="flex flex-col gap-4">
	<input type="file" name="image" accept="image/png, image/jpeg" on:change={onImageLoad} />
	<input bind:value={username} type="text" name="username" />
	<input bind:value={password} type="text" name="password" />
	<textarea bind:value={description} name="description" maxlength="150" />
	<button type="submit"> Submit </button>
	<PlayerCard {username} {description} avatar={image} />
</form>
