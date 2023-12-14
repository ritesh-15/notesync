import db from "@/lib/database/db"
import { DrizzleAdapter } from "@auth/drizzle-adapter"
import { AuthOptions } from "next-auth"

export const authOptions = {
  adapter: DrizzleAdapter(db),
  providers: [],
} satisfies AuthOptions
