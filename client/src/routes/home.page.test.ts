import { describe, jest } from "@jest/globals";
import { render, screen } from "@testing-library/svelte";
import Page from "./+page.svelte";


describe("testing the home page", () => {
    it("should load and render all agents on the page", () => {
        render(Page, {
            props: {
                data: {
                    agents: [
                        { uuid: "", displayName: "a", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                        { uuid: "", displayName: "b", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                        { uuid: "", displayName: "c", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                    ]
                }
            }
        })
        expect(screen.getAllByTestId("agent-a")).toHaveLength(1)
        expect(screen.getAllByTestId("agent-b")).toHaveLength(1)
        expect(screen.getAllByTestId("agent-c")).toHaveLength(1)

    })
})