import type Nil from "$lib/types/nil"
import { cookie } from "$lib/utils/cookie"
import server from "../../../lib/libs/axios"
import type { iPlayer } from "../../../lib/types/player"

type PlayerToCreate = iPlayer & { password: string }

const player = {

    create: async (player: PlayerToCreate) => {
        try {
            const response = await server.post("/players/register", { player }, {   
                withCredentials: true,
            })
            return response?.data.data as { id: string } | undefined
        } catch (error) {
            console.error(error)
        }
    },

    generatePassword: async () => {
        try {
            const response = await server.get("/players/generate/password")
            return response.data
        } catch (error) {
            console.error(error)
        }
    },

    getUser: async ():Promise<iPlayer | Nil> => {
        try {
            const token = cookie.get("token")
            if (token) {
                const response = await server.get("/players/auth")
                console.log(response)
            } else {
                return null
            }
        } catch (error) {
            console.error(error)
        }
    }
}

export default player