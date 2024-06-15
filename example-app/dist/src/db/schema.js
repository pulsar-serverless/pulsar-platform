"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.todo = void 0;
const pg_core_1 = require("drizzle-orm/pg-core");
exports.todo = (0, pg_core_1.pgTable)("todo", {
    id: (0, pg_core_1.serial)("id").primaryKey(),
    title: (0, pg_core_1.text)("todo_title").notNull(),
    description: (0, pg_core_1.varchar)("description", { length: 256 }).notNull(),
    completed: (0, pg_core_1.boolean)("completed").default(false),
    createdAt: (0, pg_core_1.timestamp)("created_at").defaultNow(),
});
