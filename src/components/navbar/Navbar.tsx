"use client"

import { Button } from "../ui/button"
import Link from "next/link"
import getAbsoluteURL from "@/utils/getAbsoluteURL"
import { usePathname } from "next/navigation"

const Navbar = async () => {
  const pathname = usePathname()

  return (
    <header className="w-full flex items-center justify-between px-4 py-2">
      <Link href="/">
        <h1 className="text-2xl font-bold">📒 Notesync</h1>
      </Link>
      {/* {user ? (
        <>
          <Popover>
            <PopoverTrigger>
              <Avatar>
                <AvatarImage src={""} />
                <AvatarFallback>{"Ritesh"}</AvatarFallback>
              </Avatar>
            </PopoverTrigger>
            <PopoverContent>
              <div className="flex gap-2 items-center pb-4 border-b">
                <Avatar>
                  <AvatarImage src={""} />
                  <AvatarFallback>{"Ritesh"}</AvatarFallback>
                </Avatar>
                <h1 className="font-bold text-lg">{"Ritesh"}</h1>
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
      )} */}

      <Link href={`/login?callback=${getAbsoluteURL(pathname)}`}>
        <Button>Login / Sign Up</Button>
      </Link>
    </header>
  )
}

export default Navbar
