import express from "express";
import HomeRouter from "./routes/home";

const app = express();

app.use(express.urlencoded({ extended: true }));
app.use(express.json());

app.use("/check", HomeRouter.configRouter());

app.listen(3000, "localhost", () => {
  console.log("Listening on PORT: 3000");
});
