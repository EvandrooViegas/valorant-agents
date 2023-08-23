import type { iPlayer } from "../../../lib/types/player"

type PlayerToCreate = iPlayer & { password: string }

const player = {

    create: async (player:PlayerToCreate) => {
        try {
            console.log(player)
            // const response = await server.post("/players/register", { player })
        } catch (error) {
            console.log(error)
        }
    }
}

export default player