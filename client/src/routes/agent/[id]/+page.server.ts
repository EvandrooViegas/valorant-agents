import valorant from '$lib/services/valorant/index.js'
import { error } from '@sveltejs/kit'

export async function load(ctx) {
    const params = ctx.params
    const agentId = params.id
    const agent = await valorant.getAgent(agentId)
    if(!agent) throw error(404)
    return agent 
}