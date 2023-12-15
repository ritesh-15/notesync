import { authOptions } from "@/app/api/auth/[...nextauth]/options"
import { getServerSession } from "next-auth"

const getUser = async () => {
  const session = await getServerSession(authOptions)
  return session?.user
}

export default getUser
