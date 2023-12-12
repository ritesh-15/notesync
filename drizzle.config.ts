import type { Config } from "drizzle-kit"
import * as dotenv from "dotenv"
dotenv.config({ path: ".env" })

if (!process.env.DATABASE_URL) {
  console.log("Databases URL not specified")
}

export default {
  schema: "./src/lib/database/schemas/*",
  out: "./migrations",
  driver: "pg",
  dbCredentials: {
    connectionString: process.env.DATABASE_URL || "",
  },
} satisfies Config
