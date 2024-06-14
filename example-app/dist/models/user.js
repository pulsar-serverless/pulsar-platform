"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.newUserSchema = void 0;
const zod_1 = require("zod");
class User {
    constructor(id, fullName) {
        this.id = id;
        this.fullName = fullName;
    }
}
exports.default = User;
exports.newUserSchema = zod_1.z.object({
    fullName: zod_1.z
        .string({
        required_error: "Fullname is required",
        invalid_type_error: "Fullname must be a string",
    })
        .min(2, { message: "Must be 2 or more characters long" })
        .max(256, { message: "Must be 256 or less characters long" }),
});
