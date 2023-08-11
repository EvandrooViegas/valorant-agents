
import type { iAgent } from "$lib/types/agents"
import axios from "axios"
import type { Filter } from "../../../routes/types"
const valorant = {
    getAgents: async (): Promise<iAgent[] | undefined> => {
        try {
            const response = await axios.get("http://127.0.0.1:8080/agents")
            return response.data as iAgent[]
        } catch (error: any) {
            console.log(error)
        }
    },
    filterAgents: async (filter:Filter): Promise<iAgent[] | undefined> => {
        try {
            const response = await axios.post("http://127.0.0.1:8080/agents", {
                filter
            })
            return response.data as iAgent[]
        } catch (error: any) {
            console.log(error)
        }
    }
}

export default valorant