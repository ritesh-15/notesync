import getUser from "@/utils/getUser"
import { Button } from "../ui/button"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import Link from "next/link"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import { LayoutDashboard } from "lucide-react"

const Navbar = async () => {
  const user = await getUser()

  return (
    <header className="w-full flex items-center justify-between px-4 py-2">
      <Link href="/">
        <h1 className="text-2xl font-bold">📒 Notesync</h1>
      </Link>
      {user ? (
        <>
          <Popover>
            <PopoverTrigger>
              <Avatar>
                <AvatarImage src={user.image || ""} />
                <AvatarFallback>{user.name}</AvatarFallback>
              </Avatar>
            </PopoverTrigger>
            <PopoverContent>
              <div className="flex gap-2 items-center pb-4 border-b">
                <Avatar>
                  <AvatarImage src={user.image || ""} />
                  <AvatarFallback>{user.name}</AvatarFallback>
                </Avatar>
                <h1 className="font-bold text-lg">{user.name}</h1>
              </div>
              <div className="flex gap-2 pt-4 items-center">
                <LayoutDashboard />
                <h1 className="">Dashboard</h1>
              </div>
            </PopoverContent>
          </Popover>
        </>
      ) : (
        <Link href="/login">
          <Button>Login / Sign Up</Button>
        </Link>
      )}
    </header>
  )
}

export default Navbar
