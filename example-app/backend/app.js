import express from "express";

/**
 * @typedef {import('express').Handler} Handler
 */

/**
 * Represents a serverless application.
 */
export class ServerlessApp {
  /**
   * Constructs a new instance of the application.
   *
   * Initializes the application and sets up necessary middleware and routing.
   *
   * @constructor
   */
  constructor() {
    this.app = express();
    this.functions = express.Router();

    this.port = 3000;

    this.app.use("/functions", this.functions);
    this.app.use(express.static("public"));
  }

  /**
   * Initializes the application by starting the server and performing necessary setup tasks.
   *
   * Starts the server on the specified port and performs additional initialization tasks,
   * such as logging information about the server startup and notifying the application
   * status to an external service.
   *
   * @returns {void}
   */
  initializeApp() {
    this.app.listen(this.port, async () => {
      this.log(
        "Info",
        `Example app listening at http://localhost:${this.port} with subdomain ${process.env.APP_ID}`
      );
      try {
        await fetch(
          `http://host.docker.internal:1323/app/status?subdomain=${process.env.APP_ID}`,
          { method: "post" }
        );
        this.log("Info", "Notified the app has start");
      } catch (err) {
        this.log("Error", "Notification failed");
      }
    });
  }

  /**
   * Registers a request handler for a specific endpoint at /functions/{endpoint}.
   *
   * @param {string} endpoint - The endpoint to handle requests for.
   * @param {Handler} handler - The function to be called when requests are made to the endpoint.
   * @returns {void}
   */
  onRequest(endpoint, handler) {
    this.functions.all(endpoint, handler);
  }

  /**
   * Logs a message with an optional log level.
   * @param {"Error" | "Warning" | "Info"} level - The log level.
   * @param {...any} author - The author of the book.
   */
  log(level = "Info", ...data) {
    const message = JSON.stringify({
      message: data.map((item) => item.toString()).join("\n"),
      type: level
    });

    if (level == "Error") return console.error(message);
    return console.log(message);
  }
}
