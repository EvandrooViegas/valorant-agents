import server from "../../../lib/libs/axios"

type CommentToCreate = { agent_id: string, text: string, parent_id: string }
const create = async (player: CommentToCreate) => {
    try {
        const response = await server.post("/comments/create", { player }, {
            withCredentials: true,
        })
        return response?.data.data as { id: string } | undefined
    } catch (error) {
        console.error(error)
    }
}

const comment = { create }

export default comment