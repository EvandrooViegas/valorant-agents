
import type { iAgent } from "$lib/types/agents"
import axios from "axios"
import type { Filter } from "../../../routes/types"
import type Nil from "$lib/types/nil"
import wait from "$lib/utils/wait"
const valorant = {
    getAgents: async (): Promise<iAgent[] | undefined> => {
        try {
            const response = await axios.get("http://127.0.0.1:8080/agents")
            return response.data as iAgent[]
        } catch (error: any) {
            console.log(error)
        }
    },
    getAgent: async (agentId: string): Promise<iAgent | Nil> => {
        try {
            const response = await axios.get(`http://127.0.0.1:8080/agents/${agentId}`)
            await wait()
            return response.data as iAgent
        } catch (error) {
            console.log(error)

        }
    },
    filterAgents: async (filter: Filter): Promise<iAgent[] | undefined> => {
        try {
            const response = await axios.post("http://127.0.0.1:8080/agents", {
                filter
            })
            return response.data as iAgent[]
        } catch (error: any) {
            console.log(error)
        }
    },
}

export default valorant