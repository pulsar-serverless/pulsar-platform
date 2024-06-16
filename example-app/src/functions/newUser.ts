import { PrismaClient } from "@prisma/client";
import User, { NewUser, newUserSchema } from "../models/user";
import { z } from "zod";

export default class PostUserController {
  constructor(private readonly client: PrismaClient) {}

  async postUser(data: NewUser): Promise<User> {
    const parsed = newUserSchema.safeParse(data);

    if (!parsed.success) {
      return Promise.reject(
        `Invalid data provided: ${parsed.error.issues[0].message}`
      );
    }

    try {
      const query = await this.client.user.create({
        data: parsed.data,
      });

      return new User(query.id, query.fullName);
    } catch (error) {
      return Promise.reject("error creating user");
    }
  }
}
