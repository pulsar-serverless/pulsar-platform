import { z } from "zod";

export default class User {
  constructor(public readonly id: string, public readonly fullName: string) {}
}

export const newUserSchema = z.object({
  fullName: z
    .string({
      required_error: "Fullname is required",
      invalid_type_error: "Fullname must be a string",
    })
    .min(2, { message: "Must be 2 or more characters long" })
    .max(256, { message: "Must be 256 or less characters long" }),
});

export type NewUser = z.infer<typeof newUserSchema>;
