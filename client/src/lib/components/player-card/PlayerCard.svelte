<script lang="ts">
    import defaultAvatar from "../../assets/default-avatar.png"
	import type Nil from "../../types/nil";

    export let username:Nil | string
    export let description: Nil | string
    export let avatar : Nil | string
    
</script>

<div class="w-full max-w-[700px] flex gap-4 items-center   border border-dashed border-neutral-600 p-7">
    <img src={avatar || defaultAvatar} class="w-48 aspect-square object-center object-cover" alt={`${username || "Unknow"}'s Avatar'`} />  
    <div class="flex flex-col gap-1">
        <span class="text-xl font-museo font-black bg-primary/80 px-4">@{username || "Unkow"}</span>
        <p class=" text-neutral-300 bg-neutral-00 break-all max-h-[100px] overflow-hidden">{description || "No Description"}</p>
    </div>
</div>