import type z from "zod"

type ValidateSchemaReturn = 
{ type: "success" }
| { type: "error", errors: Map<string, string> }

export default function validateSchema<T>
(schema: z.ZodObject<{}>, value: Record<string, unknown>)
: ValidateSchemaReturn {
    const errorsMap = new Map<string, string>()
    const result = schema.safeParse(value) as z.SafeParseReturnType<T, T>
    const hasErrors = 'error' in result;
    if (!hasErrors) {
        return { type: "success" }
    }

    const fields = Object.keys(value) 
    const errors = result.error.format() as z.ZodFormattedError<T, string>
    fields.forEach((field) => {
        // @ts-ignore
        const error = errors[field as keyof typeof value]?._errors[0];
        if (error) {
            errorsMap.set(field, error);
        }
    });
    return { type: "error", errors: errorsMap }
}