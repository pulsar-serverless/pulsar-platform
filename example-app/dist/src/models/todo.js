"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.newTodoSchema = void 0;
const zod_1 = __importDefault(require("zod"));
class Todo {
    constructor(id, title, description, completed, createdAt) {
        this.id = id;
        this.title = title;
        this.description = description;
        this.completed = completed;
        this.createdAt = createdAt;
    }
}
exports.default = Todo;
exports.newTodoSchema = zod_1.default.object({
    userId: zod_1.default
        .string({
        required_error: "UserID is required",
        invalid_type_error: "UserID must be a string",
    })
        .length(36, { message: "User ID should be 36 charaters long" }),
    title: zod_1.default
        .string({
        required_error: "Title is required",
        invalid_type_error: "Title must be a string",
    })
        .min(2, { message: "Must be 2 or more characters long" })
        .max(32, { message: "Must be 32 or less characters long" }),
    description: zod_1.default
        .string({
        required_error: "Description is required",
        invalid_type_error: "Description must be a string",
    })
        .min(2, { message: "Must be 2 or more characters long" })
        .max(256, { message: "Must be 256 or less characters long" }),
});
