import type { HTMLInputTypeAttribute } from "svelte/elements";

export type InputType = "textarea" | HTMLInputTypeAttribute
type Value = 
{  type: "file", url: string, value: string }
| { type: string, value: string }
export type InputContext = {
    type: string;
    name: string;
    className: string;
    label: string
    updateValue: (nValue: Value, name: string) => void
    setIsFileUploaded: (bool: boolean) => void
};   