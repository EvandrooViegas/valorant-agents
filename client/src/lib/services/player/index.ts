import type Nil from "../../../lib/types/nil"
import { cookie } from "../../../lib/utils/cookie"
import server from "../../../lib/libs/axios"
import type { iPlayer } from "../../../lib/types/player"

type PlayerToCreate = iPlayer & { password: string }
type PlayerToLogin = Pick<iPlayer, "username"> & { password: string }
const create = async (player: PlayerToCreate) => {
    try {
        const response = await server.post("/players/register", { player }, {
            withCredentials: true,
        })
        return response?.data.data as { id: string } | undefined
    } catch (error) {
        console.error(error)
    }
}

const login = async (player:PlayerToLogin) => {
    try {
        const response = await server.post("/players/login", { player }, {
            withCredentials: true,
        })
        console.log(response)
    } catch (error) {
        console.error(error)
    }
}
const getPlayer = async (): Promise<iPlayer | Nil> => {
    try {
        const token = cookie.get("token")
        if (token) {
            const response = await server.get("/players/auth", {
                withCredentials: true
            })
            return response.data?.data?.player
        } else {
            return null
        }
    } catch (error) {
        console.error(error)
    }
}
const generatePassword = async () => {
    try {
        const response = await server.get("/players/generate/password")
        return response.data
    } catch (error) {
        console.error(error)
    }
}
const player = { create, generatePassword, getPlayer, login }

export default player