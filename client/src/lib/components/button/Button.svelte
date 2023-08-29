<script lang="ts">
    import type { HTMLAttributes, HTMLButtonAttributes } from "svelte/elements";
    import { cva, type VariantProps } from "class-variance-authority";
  
    const button = cva("button", {
      variants: {
        intent: {
          primary: "bg-primary",
          border: "bg-transparent border border-dashed border-neutral-500 font-mono w-fit",
          underline: "bg-transparent underline hover:text-primary",
        },
        size: {
          small: "px-2 py-1",
          medium: "px-4 py-1.5",
        },
      },
      compoundVariants: [
        { intent: "primary", size: "medium", class: "font-bold font-mono w-fit transition-all" },
      ],
    });
  
    interface $$Props extends HTMLButtonAttributes, VariantProps<typeof button> { 
      dataTestId?: string
     }
      export let intent: $$Props["intent"] = "primary";
      export let size: $$Props["size"] = "medium";
      export let dataTestId: $$Props["dataTestId"] = ""
  </script>
  
  <button type="button" {...$$props} class={button({ intent, size, class: $$props.class })} data-testid={dataTestId}>
    <slot />
  </button>
  