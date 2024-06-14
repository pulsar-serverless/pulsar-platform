"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = require("express");
class UserRouter {
    constructor(postController) {
        this.postController = postController;
    }
    configRouter() {
        const router = (0, express_1.Router)();
        router.post("/", (req, res) => __awaiter(this, void 0, void 0, function* () {
            const newUser = req.body;
            let err;
            const response = yield this.postController
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
        }));
        return router;
    }
}
exports.default = UserRouter;
