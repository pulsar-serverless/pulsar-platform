import { ServerlessApp } from "./app.js";
import { PrismaClient } from "@prisma/client";

const myApp = new ServerlessApp();

myApp.initializeApp();

const prisma = new PrismaClient();

myApp.onRequest("/test", (req, res) => {
  return res.send(`<h1>Hello, world!</h1>`);
});

myApp.onRequest("/todos", async (req, res) => {
  for (const key in process.env) {
    myApp.log("Error",`${key}: ${process.env[key]}`);
    myApp.log("Warning",`${key}: ${process.env[key]}`);
    myApp.log("Info",`${key}: ${process.env[key]}`);
  }
  try {
    if (req.method == "GET") {
      const todos = await prisma.todo.findMany({});
      return res.status(200).json(todos);
    } else if (req.method == "POST") {
      console.log("before");
      const saved = await prisma.todo.create({
        data: {
          title: req.query.title,
          description: req.query.description,
        },
      });
      console.log("after");
      return res.status(201).json(saved);
    }
    return res.sendStatus(200);
  } catch (err) {
    console.error(err)
    return res.sendStatus(500);
  }
});
