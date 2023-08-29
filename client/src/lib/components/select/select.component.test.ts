import { describe, it, jest } from "@jest/globals"
import { render, screen, fireEvent } from "@testing-library/svelte"
import Select from "./Select.svelte"


const onSelectChangeMock = jest.fn(() => { })
describe("testing select component", () => {
    it("should display all of the options, toggle them correctly and call onSelectChange when something changes", async () => {
        render(Select, {
            props: {
                options: [
                    { id: 1, name: "a", value: "a", image: "image" },
                    { id: 2, name: "b", value: "b", image: "image" },
                    { id: 3, name: "c", value: "c", image: "image" },
                ],
                onSelectChange: onSelectChangeMock,
                selectedOptions: null,
                dataTestId: "test",
            }
        })

        const selectComp = screen.getByTestId("test")
        await fireEvent.click(selectComp)

        const optA = screen.getByTestId(`select-option-a`)
        await fireEvent.click(optA)
        expect(optA.id).toBe("on")
        await fireEvent.click(optA)
        expect(optA.id).toBe("off")

        const optB = screen.getByTestId(`select-option-b`)
        await fireEvent.click(optB)
        expect(optB.id).toBe("on")
        await fireEvent.click(optB)
        expect(optB.id).toBe("off")
        expect(onSelectChangeMock).toBeCalled()

        const optC = screen.getByTestId(`select-option-c`)
        await fireEvent.click(optC)
        expect(optC.id).toBe("on")
        await fireEvent.click(optC)
        expect(optC.id).toBe("off")
        expect(onSelectChangeMock).toBeCalled()

        expect(optA).toBeTruthy()
        expect(optB).toBeTruthy()
        expect(optC).toBeTruthy()
        expect(onSelectChangeMock).toBeCalledTimes(6)
    })

})