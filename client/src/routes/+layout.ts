import type { iPlayer } from "$lib/types/player"
import player from "../lib/services/player"

export async function load() {
    try {
        const p = await player.getPlayer() as iPlayer | undefined
        return { player: p, isLoggedIn: p ? true : false } 
    } catch (error) {
        console.log(error)
    }
}