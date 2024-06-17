import { ServerlessApp } from "./app.js";

const myApp = new ServerlessApp();

myApp.initializeApp();

myApp.onRequest("/", (req, res) => {
  return res.send(`<h1>Hello, world!</h1>`);
});
