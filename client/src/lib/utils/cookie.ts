export const cookie = {
    getAll: function():Map<string, string> | null  {
        const cookies = document.cookie?.split?.(";").map?.(c => c?.split("="))
        if (cookies) {
            return  new Map<string, string>()
        } else {
            return null
        }
    },
    get: function(name: string):string | undefined  {
        const cookies = this.getAll()
        if (cookies) {
            cookies.get(name)
        } else {
            return undefined
        }
    }
}