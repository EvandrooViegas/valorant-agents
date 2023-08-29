import { describe, it, jest } from "@jest/globals"
import { render, screen, fireEvent } from "@testing-library/svelte"
import Select from "./Select.svelte"


const onSelectChangeMock = jest.fn(() => { })
describe("testing select component", () => {
    it("should display all of the options", async () => {
        render(Select, {
            props: {
                options: [
                    { id: 1, name: "a", value: "a", image: "image" },
                    { id: 2, name: "b", value: "b", image: "image" },
                    { id: 3, name: "c", value: "c", image: "image" },
                ] ,
                onSelectChange: onSelectChangeMock,
                selectedOptions: null,
                dataTestId: "test",
            }
        })

        const selectComp = screen.getByTestId("test")
        await fireEvent.click(selectComp)
        expect(screen.getAllByTestId(`select-option-a`)).toHaveLength(1)
        expect(screen.getAllByTestId(`select-option-b`)).toHaveLength(1)
        expect(screen.getAllByTestId(`select-option-c`)).toHaveLength(1)
    })
})