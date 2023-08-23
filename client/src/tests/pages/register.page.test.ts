
import { describe, it, expect } from "@jest/globals"
import { render, screen } from '@testing-library/svelte'
import Page from "../../routes/auth/register/+page.svelte"

describe("Testing register page", () => {
    it("should call the function to create the player if the field are filled correctly", () => {
        render(Page)
        const spanEl = screen.queryAllByText('Hello World Evandro')
        expect(spanEl.length).toBe(1)
    })
})