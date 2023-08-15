export const actions = {
    create: async ({ request }) => {
        const form = await request.formData()
        console.log(form)
    }
}