import { Request, Response, Router } from "express";
import GetTodoController from "../functions/getTodo";
import PostTodoController from "../functions/newTodo";

export default class TodoRouter {
  constructor(
    private readonly getController: GetTodoController,
    private readonly postController: PostTodoController
  ) {}

  configRouter() {
    const router = Router();

    router.get("/", async (req: Request, res: Response) => {
      const userID = req.query.id as string;
      const todoID = parseInt(req.query.todoID as string);

      let err: any;

      const response = await this.getController
        .getTodo(todoID, userID)
        .catch((error) => {
          err = error;
          return null;
        });

      if (!response) {
        return res.status(400).json({
          msg: err,
        });
      }

      return res.status(200).json({
        data: response,
      });
    });

    router.get("/all", async (req: Request, res: Response) => {
      const userID = req.query.id as string;
      console.log(userID)

      let err: any;

      const response = await this.getController
        .getTodos(userID)
        .catch((error) => {
          err = error;
          return null;
        });

      if (!response) {
        return res.status(400).json({
          msg: err,
        });
      }

      return res.status(200).json({
        data: response,
      });
    });

    router.post("/", async (req: Request, res: Response) => {
      const userID = req.query.id as string
      const newTodo = req.body;

      let err: any;

      const response = await this.postController
        .postTodo({ ...newTodo, userId: userID })
        .catch((error) => {
          err = error;
          return null;
        });

      if (!response) {
        return res.status(400).json({
          msg: err,
        });
      }

      return res.status(201).json({
        data: response,
      });
    });

    return router;
  }
}
