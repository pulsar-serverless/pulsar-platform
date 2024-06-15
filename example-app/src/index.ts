import express from "express";
import HomeRouter from "./routes/home";
import TodoRouter from "./routes/todo.router";
import dotenv from "dotenv";
import { PrismaClient } from "@prisma/client";
import GetTodoController from "./functions/getTodo";
import PostTodoController from "./functions/newTodo";
import PostUserController from "./functions/newUser";
import UserRouter from "./routes/user.router";

dotenv.config({ path: "../.env" });

const client = new PrismaClient();
const getController = new GetTodoController(client);
const postController = new PostTodoController(client);
const userPostContoller = new PostUserController(client);

const userRouter = new UserRouter(userPostContoller);
const todoRouter = new TodoRouter(getController, postController);

const PORT = parseInt(process.env.PORT!);
const app = express();

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

app.use("/api/v1/check", HomeRouter.configRouter());
app.use("/api/v1/user", userRouter.configRouter());
app.use("/api/v1/todo", todoRouter.configRouter());

app.listen(PORT, "localhost", () => {
  console.log("Listening on PORT: ", PORT);
});
