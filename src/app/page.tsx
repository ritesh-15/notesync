import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { getServerSession } from "next-auth"
import { redirect } from "next/navigation"
import { authOptions } from "./api/auth/[...nextauth]/options"

const Home = async () => {
  const session = await getServerSession(authOptions)

  if (!session) redirect("/login")

  return (
    <main className="">
      <h1>Hello</h1>
    </main>
  )
}

export default Home
