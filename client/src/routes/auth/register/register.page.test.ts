import { describe, it, expect, jest } from "@jest/globals";
import { render, fireEvent, screen } from "@testing-library/svelte";
import Page from "./+page.svelte";
import player from "../../../lib/services/player";
import storage from "../../../lib/services/storage";
describe("register page", () => {
    global.window.URL.createObjectURL= jest.fn() as jest.Mock<() => string>
    it("should render all of the inputs correctly", async () => {
        render(Page);
        // get the inputs
        const usernameInput = screen.getByTestId("username");
        const passwordInput = screen.getByTestId("password") as HTMLInputElement;
        const descriptionInput = screen.getByTestId("description") as HTMLInputElement;
        const avatarInput = screen.getByTestId("avatar") as HTMLInputElement;

        expect(usernameInput).toBeTruthy()
        expect(passwordInput).toBeTruthy()
        expect(descriptionInput).toBeTruthy()
        expect(avatarInput).toBeTruthy()
    })
    it("should call the function to create the player if the field are filled correctly", async () => {
        const playerCreateMock = jest.spyOn(player, "create")
        const storageUploadMock = jest.spyOn(storage, "uploadImage")
        playerCreateMock.mockImplementation(() => Promise.resolve())
        storageUploadMock.mockImplementation(() => Promise.resolve({ url: "hello", name: "url", size: 13 }))
        // render the page
        render(Page);
        // get the inputs and form
        const form = screen.getByTestId("form") as HTMLFormElement
        const usernameInput = screen.getByTestId("username") as HTMLInputElement
        const passwordInput = screen.getByTestId("password") as HTMLInputElement;
        const descriptionInput = screen.getByTestId("description") as HTMLInputElement;
        const avatarInput = screen.getByTestId("avatar") as HTMLInputElement;
        
        const file = new File(["hello world"], "file.txt")
        // fill of the fields correctly
        await fireEvent.input(usernameInput, { target: { value: "Evandro" } });
        await fireEvent.input(passwordInput, { target: { value: "mypassword12345" } });
        await fireEvent.input(descriptionInput, { target: { value: "My description" } });
        await fireEvent.change(avatarInput, { target: { files: [file] } });

        // submit
        form.submit()

        // expect that the player.create function was called
        expect(player.create).toBeCalled();
        playerCreateMock.mockRestore()
        storageUploadMock.mockRestore()

    });
    it("should not call the function to create the player if the field are filled correctly and show up error messages", async () => {
        const playerCreateMock = jest.spyOn(player, "create")
        const storageUploadMock = jest.spyOn(storage, "uploadImage")
        playerCreateMock.mockImplementation(() => Promise.resolve())
        storageUploadMock.mockImplementation(() => Promise.resolve({ url: "hello", name: "url", size: 13 }))
        // render the page
        render(Page);
        // get the inputs and form
        const form = screen.getByTestId("form") as HTMLFormElement
        const usernameInput = screen.getByTestId("username") as HTMLInputElement
        const passwordInput = screen.getByTestId("password") as HTMLInputElement;
        const descriptionInput = screen.getByTestId("description") as HTMLInputElement;
        
        // fill of the fields correctly
        await fireEvent.input(usernameInput, { target: { value: "" } }); // error
        await fireEvent.input(passwordInput, { target: { value: "" } }); // error 
        await fireEvent.input(descriptionInput, { target: { value: "" } }); // error
        
        // expect that the errors messages are being not shown before submiting the form 
        expect(screen.queryAllByTestId("username-error")).toHaveLength(0)
        expect(screen.queryAllByTestId("password-error")).toHaveLength(0)
        expect(screen.queryAllByTestId("description-error")).toHaveLength(0)
        
        // submit
        await form.submit()

        // expect that the errors messages are being shown after submiting the form
        expect(screen.queryAllByTestId("username-error")).toHaveLength(1)
        expect(screen.queryAllByTestId("password-error")).toHaveLength(1)
        expect(screen.queryAllByTestId("description-error")).toHaveLength(1)
        // expect that the player.create function was called
        expect(player.create).not.toBeCalled();
        playerCreateMock.mockRestore()
        storageUploadMock.mockRestore()
    });
})