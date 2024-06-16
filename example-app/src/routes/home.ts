import { Request, Response, Router } from "express";

export default class HomeRouter {
  static configRouter() {
    const router = Router();

    router.get("/", async (req: Request, res: Response) => {
      return res.status(200).json({
        msg: "Server running...",
      });
    });

    return router;
  }
}
