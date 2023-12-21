import { drizzle } from "drizzle-orm/postgres-js"
import { migrate } from "drizzle-orm/postgres-js/migrator"
import postgres from "postgres"
import * as dotenv from "dotenv"
import * as schemas from "../database/schemas/index"
dotenv.config({ path: ".env" })

if (!process.env.DATABASE_URL) {
  console.log("Database URL not specified")
}

const client = postgres(process.env.DATABASE_URL as string, {
  keep_alive: 1,
  max: 300,
})
const db = drizzle(client, { schema: schemas })

const migrateDb = async () => {
  try {
    console.log("Migrating client...")
    await migrate(db, { migrationsFolder: "migrations" })
    console.log("Migration successful")
  } catch (err) {
    console.log("Failed to migrate client")
    console.error(err)
  }
}

// migrateDb()
export default db
