import player from "$lib/services/player"

export async function load(ctx) {
    const params = ctx.params
    const username = params.username
    const p = await player.getPlayerByUsername(username)
    return { player: p, wasPlayerFound: !!p  } 
}

