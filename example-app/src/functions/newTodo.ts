import { PrismaClient } from "@prisma/client";
import Todo, { NewTodo, newTodoSchema } from "../models/todo";

export default class PostTodoController {
  constructor(private readonly client: PrismaClient) {}

  async postTodo(data: NewTodo): Promise<Todo> {
    const parsed = newTodoSchema.safeParse(data);

    if (!parsed.success) {
      return Promise.reject(
        `Invalid data provided: ${parsed.error.issues[0].message}`
      );
    }

    try {
      const query = await this.client.todo.create({
        data: parsed.data,
      });

      return new Todo(
        query.id,
        query.title,
        query.description,
        query.completed,
        query.createdAt
      );
    } catch (error) {
      return Promise.reject("error creating todo");
    }
  }
}
