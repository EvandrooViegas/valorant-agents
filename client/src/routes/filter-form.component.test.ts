import { describe, expect, it, jest } from "@jest/globals";
import { render, screen, fireEvent } from "@testing-library/svelte";
import FilterForm from "./FilterForm.svelte";

const onSubmitMock = jest.fn(() => Promise.resolve())
const resetFiltersMock = () => { }
describe("Testing the filter form component from the home page", () => {
    it("shouldn't show the filter and remove filters button if the input is empty", () => {
        render(FilterForm, {
            props: {
                agents: [
                    { uuid: "", displayName: "a", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                    { uuid: "", displayName: "b", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                    { uuid: "", displayName: "c", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                ],
                name: "",
                onSubmit: onSubmitMock,
                resetFilteredAgents: resetFiltersMock,
            }
        })

        expect(screen.queryAllByTestId("filter-btn")).toHaveLength(0)
        expect(screen.queryAllByTestId("remove-filter-btn")).toHaveLength(0)
    })

    it("should show the filter and remove filters button when the text input has some value", () => {
        render(FilterForm, {
            props: {
                agents: [
                    { uuid: "", displayName: "a", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                    { uuid: "", displayName: "b", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                    { uuid: "", displayName: "c", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                ],
                name: "Person",
                onSubmit: onSubmitMock,
                resetFilteredAgents: resetFiltersMock,
            }
        })

        expect(screen.queryAllByTestId("filter-btn")).toHaveLength(1)
        expect(screen.queryAllByTestId("remove-filter-btn")).toHaveLength(1)
    })
    it("should call the onSubmit function when the button is clicked", async () => {
        render(FilterForm, {
            props: {
                agents: [
                    { uuid: "", displayName: "a", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                    { uuid: "", displayName: "b", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                    { uuid: "", displayName: "c", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "", displayIcon: "" }, abilities: [] },
                ],
                name: "Person",
                onSubmit: onSubmitMock,
                resetFilteredAgents: resetFiltersMock,
            }
        })
        const submitButton = screen.getByTestId("filter-btn")
        await fireEvent.submit(submitButton)
        expect(onSubmitMock).toBeCalled()
    })
    it("should remove all of the filters when the remove filters button is clicked", async () => {
        render(FilterForm, {
            props: {
                agents: [
                    { uuid: "", displayName: "a", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "a", displayIcon: "" }, abilities: [] },
                    { uuid: "", displayName: "b", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "b", displayIcon: "" }, abilities: [] },
                    { uuid: "", displayName: "c", description: "", displayIcon: "", fullPortrait: "", background: "", role: { displayName: "c", displayIcon: "" }, abilities: [] },
                ],
                name: "Hello",
                onSubmit: onSubmitMock,
                resetFilteredAgents: resetFiltersMock,
            }
        })

        const removeFilterBtn = screen.getByTestId("remove-filter-btn")
        await fireEvent.click(removeFilterBtn)
        const textInput = screen.getByTestId("filter-form-name-input") as HTMLInputElement
        expect(textInput.value).toBe("")
    })
})