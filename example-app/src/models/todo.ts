import z from "zod";

export default class Todo {
  constructor(
    public readonly id: number,
    public readonly title: string,
    public readonly description: string,
    public readonly completed: boolean,
    public readonly createdAt: Date
  ) {}
}

export const newTodoSchema = z.object({
  userId: z
    .string({
      required_error: "UserID is required",
      invalid_type_error: "UserID must be a string",
    })
    .length(36, { message: "User ID should be 36 charaters long" }),
  title: z
    .string({
      required_error: "Title is required",
      invalid_type_error: "Title must be a string",
    })
    .min(2, { message: "Must be 2 or more characters long" })
    .max(32, { message: "Must be 32 or less characters long" }),
  description: z
    .string({
      required_error: "Description is required",
      invalid_type_error: "Description must be a string",
    })
    .min(2, { message: "Must be 2 or more characters long" })
    .max(256, { message: "Must be 256 or less characters long" }),
});

export type NewTodo = z.infer<typeof newTodoSchema>;
