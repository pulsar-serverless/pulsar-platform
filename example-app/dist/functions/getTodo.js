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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const todo_1 = __importDefault(require("../models/todo"));
class GetTodoController {
    constructor(client) {
        this.client = client;
    }
    getTodo(id, userId) {
        return __awaiter(this, void 0, void 0, function* () {
            if (!userId) {
                return Promise.reject("user id not provided");
            }
            if (!id) {
                return Promise.reject("todo id not provided");
            }
            const query = yield this.client.todo.findUnique({
                where: {
                    id: id,
                    userId: userId,
                },
            });
            if (!query) {
                return Promise.reject("Todo not found!");
            }
            return new todo_1.default(query.id, query.title, query.description, query.completed, query.createdAt);
        });
    }
    getTodos(userId) {
        return __awaiter(this, void 0, void 0, function* () {
            if (!userId) {
                return Promise.reject("user id not provided");
            }
            const query = yield this.client.todo.findMany({
                where: {
                    userId: userId,
                },
            });
            if (query.length <= 0) {
                return Promise.reject("Todos not found!");
            }
            return query.map((val) => {
                return new todo_1.default(val.id, val.title, val.description, val.completed, val.createdAt);
            });
        });
    }
}
exports.default = GetTodoController;
