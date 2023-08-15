<script lang="ts">
	import Button from '$components/button/Button.svelte';
	import Input from '$components/input/Input.svelte';
	import PlayerCard from '$components/player-card/PlayerCard.svelte';

	let image: string;
	let username: string;
	let password: string;
	let description: string;
</script>

<form method="POST" action="?/create" enctype="multipart/form-data" class="grid grid-cols-2 gap-4">
	<Input label="Username" bind:value={username} name="username" />
	<Input label="Password" bind:value={password} name="password" />
	<Input
		label="Description"
		bind:value={description}
		name="description"
		maxlength="150"
		type="textarea"
		containerClassName="col-span-2"
	/>
	<Input
		label="Avatar"
		bind:value={image}
		name="image"
		type="file"
		accept="image/png, image/jpeg"
	/>
	<Button type="submit" class="col-span-2">Submit</Button>

    <div>
        <span>Preview: </span>
        <PlayerCard {username} {description} avatar={image} />
    </div>
</form>
