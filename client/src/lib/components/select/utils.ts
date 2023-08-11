import type { iAgent, iRole } from "$lib/types/agents";
import type { Option } from "./type";

export const convertOptionsToRoles = (options: Option<iRole>[] | undefined | null): iRole[] => {
    if (!options) return []
    return options.map(opt => opt.value)
}

export const converRolesToOptions = (roles: iRole[] | undefined): Option<iRole>[] => {
    if (!roles) return []
    return roles?.map((r, idx) => ({
        id: idx,
        name: r.displayName,
        value: r,
        image: r.displayIcon
    }))
}

export const getRolesFromAgents = (agents: iAgent[] | undefined) => {
    return agents
        ?.reduce((curr, agent) => {
            const exists = curr.map((c) => c.displayName)
                .includes(agent.role.displayName);
            return exists ? [...curr] : [...curr, agent.role];
        }, [] as iRole[])
        .filter((r) => r.displayName !== '' || r.displayIcon !== '') 

}