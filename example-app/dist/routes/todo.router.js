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
class TodoRouter {
    constructor(getController, postController) {
        this.getController = getController;
        this.postController = postController;
    }
    configRouter() {
        const router = (0, express_1.Router)();
        router.get("/", (req, res) => __awaiter(this, void 0, void 0, function* () {
            const userID = req.query.id;
            const todoID = parseInt(req.query.todoID);
            let err;
            const response = yield this.getController
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
        }));
        router.get("/all", (req, res) => __awaiter(this, void 0, void 0, function* () {
            const userID = req.query.id;
            console.log(userID);
            let err;
            const response = yield this.getController
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
        }));
        router.post("/", (req, res) => __awaiter(this, void 0, void 0, function* () {
            const userID = req.query.id;
            const newTodo = req.body;
            let err;
            const response = yield this.postController
                .postTodo(Object.assign(Object.assign({}, newTodo), { userId: userID }))
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
exports.default = TodoRouter;
