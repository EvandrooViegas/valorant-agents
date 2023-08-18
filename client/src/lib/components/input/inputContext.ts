import { createContext } from "$lib/utils/createContext"

export function inputContext<T = any>(value?: T) {  
    const context = createContext<T>("input")
	context.set(value || "")
    return context.get()
}