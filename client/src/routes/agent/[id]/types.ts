import type { iAgent } from "$lib/types/agents"
import type z from "zod"

export type CommentFormContext = {
    formSchema: z.ZodObject<{}> 
}

export type AgentContext = {
    agent: iAgent
}

export type InputContextContent = {
    formSchema: z.ZodObject<{}> 
}