import { defineConfig } from "cypress";

export default defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },

  env: {
    auth0_domain: process.env.NEXT_PUBLIC_AUTH0_DOMAIN,
    auth0_audience: process.env.NEXT_PUBLIC_AUTH0_AUDIENCE,
    auth0_client_id: process.env.NEXT_PUBLIC_AUTH0_CLIENT_ID,
    auth0_client_secret: process.env.NEXT_PUBLIC_AUTH0_CLIENT_SECRETE,
    auth0_redirect_uri: process.env.NEXT_PUBLIC_REDIRECT_URI,
  },

  component: {
    devServer: {
      framework: "next",
      bundler: "webpack",
    },
  },
});
