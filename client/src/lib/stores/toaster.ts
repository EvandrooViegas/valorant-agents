import { writable, type Updater } from "svelte/store";

type Payload = {
    title: string
    description?: string
    type?: "success" | "error"
}

export type Item = { id: number } & Payload
type StoreUpdateFunc = (this: void, updater: Updater<Item[]>) => void

const toastVisibilityDuration = 1 * 3000 // 3 seconds

const add = (payload: Payload, update: StoreUpdateFunc) => {
    const toastId = Math.random()
    const newItem:Item = { id: toastId, ...payload }
    update((toasts) => {
        return [...toasts, newItem]
    })

    setTimeout(() => {
        // remove item from toast arr
        update((toasts) => {
            return toasts.filter(item => item.id !== toastId)
        })
    }, toastVisibilityDuration)
}

function createToaster() {
    const { subscribe, update } = writable<Item[]>([])
    return {
        subscribe,
        add: (payload:Payload) => add(payload, update)
    }
}

export const toaster = createToaster()