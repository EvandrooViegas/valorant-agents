import { describe, jest } from "@jest/globals";
import { fireEvent, render, screen } from "@testing-library/svelte";
import Page from "./+page.svelte";
import valorant from "../lib/services/valorant";
import type { iRole } from "../lib/types/agents";

const agents = [
    { uuid: "", displayName: "a", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
    { uuid: "", displayName: "b", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
    { uuid: "", displayName: "c", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
]
describe("testing the home page", () => {
    it("should load and render all agents on the page", () => {
        render(Page, {
            props: { data: { agents } }
        })
        expect(screen.getAllByTestId("agent-a")).toHaveLength(1)
        expect(screen.getAllByTestId("agent-b")).toHaveLength(1)
        expect(screen.getAllByTestId("agent-c")).toHaveLength(1)
    })

    it("should filter the agents and display them", async () => {
        jest.spyOn(valorant, "filterAgents")
            .mockImplementation((filter: { name: string, roles: iRole[] | undefined }) => {
                const filteredAgents = agents.filter(a => a.displayName.includes(filter.name))
                return Promise.resolve(filteredAgents)
            })
        render(Page, {
            props: { data: { agents } }
        })
        expect(screen.getAllByTestId("agent-a")).toHaveLength(1)
        expect(screen.getAllByTestId("agent-b")).toHaveLength(1)
        expect(screen.getAllByTestId("agent-c")).toHaveLength(1)

        const nameInput = screen.getByTestId("filter-form-name-input") as HTMLInputElement
        await fireEvent.input(nameInput, { target: { value: "a" } })
        expect(nameInput.value).toBe("a")

        const filterBtn = screen.getByTestId("filter-btn")
        await fireEvent.submit(filterBtn);
        expect(valorant.filterAgents).toBeCalledWith({ name: "a", roles: [] })

        expect(screen.queryAllByTestId("agent-a")).toHaveLength(1)
        expect(screen.queryAllByTestId("agent-b")).toHaveLength(0)
        expect(screen.queryAllByTestId("agent-c")).toHaveLength(0)

    })
})