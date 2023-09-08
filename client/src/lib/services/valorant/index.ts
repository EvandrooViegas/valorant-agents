
import type { iAgent } from "../../../lib/types/agents"
import type { Filter } from "../../../routes/types"
import type Nil from "../../../lib/types/nil"
import server from "../../../lib/libs/axios"
const valorant = {
    getAgents: async (): Promise<iAgent[] | undefined> => {
        try {
            const response = await server.get("/agents")
            return response.data?.data as iAgent[]
        } catch (error: any) {
            console.log(error)
        }
    },
    getAgent: async (agentId: string): Promise<iAgent | Nil> => {
        try {
            const response = await server.get(`/agents/${agentId}`)
            return response.data?.data as iAgent
        } catch (error) {
            console.log(error)

        }
    },
    filterAgents: async (filter: Filter): Promise<iAgent[] | undefined> => {
        try {
            const response = await server.post("/agents/filter", {
                filter
            })
            return response.data?.data as iAgent[]
        } catch (error: any) {
            console.log(error)
        }
    },
}

export default valorant