import { PrismaClient } from "@prisma/client";
import Todo from "../models/todo";

export default class GetTodoController {
  constructor(private readonly client: PrismaClient) {}

  async getTodo(id: number, userId: string): Promise<Todo> {
    if (!userId) {
      return Promise.reject("user id not provided");
    }

    if (!id) {
      return Promise.reject("todo id not provided");
    }
    
    const query = await this.client.todo.findUnique({
      where: {
        id: id,
        userId: userId,
      },
    });

    if (!query) {
      return Promise.reject("Todo not found!");
    }

    return new Todo(
      query.id,
      query.title,
      query.description,
      query.completed,
      query.createdAt
    );
  }

  async getTodos(userId: string): Promise<Todo[]> {
    if (!userId) {
      return Promise.reject("user id not provided");
    }

    const query = await this.client.todo.findMany({
      where: {
        userId: userId,
      },
    });

    if (query.length <= 0) {
      return Promise.reject("Todos not found!");
    }

    return query.map((val) => {
      return new Todo(
        val.id,
        val.title,
        val.description,
        val.completed,
        val.createdAt
      );
    });
  }
}
