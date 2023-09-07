export const cookie = {
    getAll: function():Map<string, string> | null  {
        const cookies = document.cookie?.split?.(";").map?.(c => {
            const splitedKey = c?.split("=")[0]
            const splitedValue = c?.split("=")[1]
            const key = splitedKey.includes(" ") ? splitedKey.split(" ")[1] : splitedKey as string
            const value = splitedValue as string
            return [key, value]
        })
        if (cookies) {
            return  new Map<string, string>(cookies as Iterable<readonly [string, string]>)
        } else {
            return null
        }
    },
    get: function(name: string):string | undefined  {
        const cookies = this.getAll()
        console.log(cookies?.keys())
       return cookies?.get(name)
    }
}