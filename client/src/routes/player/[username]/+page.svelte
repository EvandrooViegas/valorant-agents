<script lang="ts">
    export let data 
    $: wasPlayerFound = data.wasPlayerFound 
    $: player = data.player
</script>

{#if wasPlayerFound}
    <div class="max-w-[1000px] mx-auto">
        <div class="relative flex justify-between items-center gap-4 p-12">
            <div class="absolute bg-gradient-to-b from-neutral-500 to-neutral-600 inset-0 h-1/2 z-[-1]">
                <div class="px-4 w-full h-full flex items-center justify-end">
                    <span class="text-4xl font-museo font-black">@{player.username}</span>

                </div>
            </div>
            
            <img src={player.avatar} alt="Player avatar" class="w-64 aspect-square object-cover rounded-full" />
        </div>
        <p class="text-sm text-neutral-400">
            {player?.description}
        </p>
    </div>
{:else}
     Couldn't find the player
{/if}