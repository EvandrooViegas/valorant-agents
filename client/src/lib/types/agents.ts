export interface iRole {
    displayName: string
    displayIcon: string
}

export interface iAbilities {
    displayName: string
    description: string
    displayIcon: string
}

export interface iAgent {
    uuid: string
    displayName: string
    description: string
    displayIcon: string
    fullPortrait: string
    background: string
    role: iRole
    abilities: iAbilities[]
}
