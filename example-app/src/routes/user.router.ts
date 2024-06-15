import { Request, Response, Router } from "express";
import PostUserController from "../functions/newUser";

export default class UserRouter {
  constructor(private readonly postController: PostUserController) {}

  configRouter() {
    const router = Router();

    router.post("/", async (req: Request, res: Response) => {
      const newUser = req.body;

      let err: any;

      const response = await this.postController
        .postUser(newUser)
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
