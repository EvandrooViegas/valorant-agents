jest.mock("./+page.server", () => {
    const originalModule = jest.requireActual("./+page.server") as {}
    return {
        __esModule: true,
        ...originalModule,
        load: jest.fn(() => ({
            agents: [
                { uuid: "", displayName: "a", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                { uuid: "", displayName: "b", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                { uuid: "", displayName: "c", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
            ]
        }))
    };
})

import { describe, jest } from "@jest/globals";
import { render, screen } from "@testing-library/svelte";
import Page from "./+page.svelte";
import * as loadModule from "./+page.server";


describe("testing the home page", () => {
    it("should load and render all agents on the page", () => {
        render(Page)
        expect(loadModule.load).toBeCalled()
        expect(screen.getAllByTestId("agent-a")).toHaveLength(1)
        expect(screen.getAllByTestId("agent-b")).toHaveLength(1)
        expect(screen.getAllByTestId("agent-c")).toHaveLength(1)

    })
})