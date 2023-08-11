import type { iRole } from "$lib/types/agents"

export interface Filter {
    name: string
    roles: iRole[]
}