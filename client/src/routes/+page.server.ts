import valorant from "../lib/services/valorant"

export async function load() {
    const agents = await valorant.getAgents()
    return { agents } 
}

