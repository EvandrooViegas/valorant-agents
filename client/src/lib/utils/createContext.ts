import { getContext, setContext } from "svelte";

export type Context<T> = {
    get: () => T
    set: (ctx: T) => void
}

export function createContext<T>(key: string):Context<T> {
    return {
        get: () => getContext<T>(key),
        set: (ctx: T) => setContext(key, ctx)
    }
}