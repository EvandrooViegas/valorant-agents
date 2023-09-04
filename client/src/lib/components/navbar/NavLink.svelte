<script lang="ts">
    export let link: { url: string, name: string }
    
</script>

<li>
    <a class="text-semibod hover:underline" href={link.url}>{link.name}</a>
</li>