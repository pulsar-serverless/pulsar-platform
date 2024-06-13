"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = require("express");
class HomeRouter {
    static configRouter() {
        const router = (0, express_1.Router)();
        router.get("/", async (req, res) => {
            return res.status(200).json({
                msg: "Server running...",
            });
        });
        return router;
    }
}
exports.default = HomeRouter;
