import valorant from '../../../lib/services/valorant/index.js'
import type { iAgent } from '../../../lib/types/agents.js'
import { error } from '@sveltejs/kit'
export async function load(ctx) {
    const params = ctx.params
    const agentId = params.id
    const agent = valorant.getAgent(agentId) as Promise<iAgent>
    if(!agentId) throw error(404)
    return { agent }
}