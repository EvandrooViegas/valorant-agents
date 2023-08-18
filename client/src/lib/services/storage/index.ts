import server from "$lib/libs/axios"
import type Nil from "$lib/types/nil"

type Image = {
    name: string,
    size: number,
    url: string
}
const storage = {
    upload: async (form:FormData):Promise<Nil | Image> => {
        try {
            const response = await server.post("/storage/upload", form)
            return response?.data?.image
        } catch (error) {
            console.log(error)
        }
    }
}

export default storage