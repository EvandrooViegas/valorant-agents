import type { iPlayer } from "./player"

export interface iComment {
    id: string
    agent_id: string
    parent_id: string
    text: string
    author: iPlayer
    replies: iComment[]
}