import type { HTMLInputTypeAttribute } from "svelte/elements";

export type InputType = "textarea" | HTMLInputTypeAttribute
export type InputContext = {
    type: string;
    name: string;
    className: string;
    label: string
    updateValue: (nValue: any, name: string) => void
};