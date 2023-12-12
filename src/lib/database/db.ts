import { drizzle } from "drizzle-orm/postgres-js"
import postgres from "postgres"
import * as dotenv from "dotenv"
import * as schemas from "../database/schemas/index"
dotenv.config()

if (!process.env.DATABASE_URL) {
  console.log("Database URL not specified")
}

const client = postgres(process.env.DATABASE_URL as string)
const db = drizzle(client, { schema: schemas })
