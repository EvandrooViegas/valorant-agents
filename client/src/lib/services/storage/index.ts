import server from "$lib/libs/axios"
import type Nil from "$lib/types/nil"

type Image = {
    name: string,
    size: number,
    url: string
}
const storage = {
    uploadImage: async (file: File, key: string = "avatar"): Promise<Nil | Image> => {
        try {
            const form = new FormData();
            form.append(key, file);
            const response = await server.post("/storage/upload", form)
            return response?.data?.image
        } catch (error) {
            console.log(error)
        }
    },
    deleteImage: async (name: string) => {
        try {
            const response = await server.delete(`/storage/delete/${name}`)
        } catch (error) {
            console.log(error)
        }
    }
}

export default storage