import db from "@/lib/database/db"
import { DrizzleAdapter } from "@auth/drizzle-adapter"
import { NextAuthOptions } from "next-auth"
import GoogleProvider from "next-auth/providers/google"
import Email from "next-auth/providers/email"

export const authOptions = {
  adapter: DrizzleAdapter(db),
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID!!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!!,
    }),
    Email({
      server: {
        host: process.env.SMTP_HOST,
        port: Number(process.env.SMTP_PORT),
        auth: {
          user: process.env.SMTP_USER,
          pass: process.env.SMTP_PASSWORD,
        },
      },
      from: process.env.EMAIL_FROM,
    }),
  ],
  pages: {
    signIn: "/login",
  },
} satisfies NextAuthOptions
