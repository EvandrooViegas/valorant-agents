import server from "../../../lib/libs/axios"
import type { iPlayer } from "../../../lib/types/player"

type PlayerToCreate = iPlayer & { password: string }

const player = {

    create: async (player: PlayerToCreate) => {
        try {
            const response = await server.post("/players/register", { player }, {   
                withCredentials: true,
            })
            console.log(response)
        } catch (error) {
            console.log(error)
        }
    },

    generatePassword: async () => {
        try {
            const response = await server.get("/players/generate/password")
            return response.data
        } catch (error) {
            console.log(error)
        }
    }
}

export default player