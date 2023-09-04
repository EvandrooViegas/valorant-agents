import player from "$lib/services/player"

export async function load() {
    const response = await player.getUser()
    return { agents: [] } 
}