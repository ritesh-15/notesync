"use client"

import { cn } from "@/lib/utils"
import {
  ChevronRight,
  ChevronsLeft,
  History,
  MenuIcon,
  PlusCircle,
  Search,
  Settings,
} from "lucide-react"
import { usePathname } from "next/navigation"
import { ElementRef, useRef, useState } from "react"
import { useMediaQuery } from "usehooks-ts"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"

const Sidebar = () => {
  const pathname = usePathname()
  const isMobile = useMediaQuery("(max-width: 768px)")

  const isResizingRef = useRef(false)
  const sidebarRef = useRef<ElementRef<"aside">>(null)
  const navbarRef = useRef<ElementRef<"div">>(null)
  const [isResetting, setIsResetting] = useState(false)
  const [isCollapsed, setIsCollapsed] = useState(false)

  return (
    <>
      <aside
        ref={sidebarRef}
        className={cn(
          "flex z-[999] h-full overflow-y-auto w-80 bg-secondary relative flex-col group/sidebar",
          isResetting && "transition-all ease-in-out duration-300",
          isMobile && "w-0"
        )}
      >
        <div className="flex items-center px-2 py-2 relative hover:bg-neutral-200">
          <div className="flex items-center gap-2 cursor-pointer ">
            <Avatar>
              <AvatarImage
                className=""
                src="https://github.com/shadcn.png"
                alt="@shadcn"
              />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>
            <h1 className=" font-bold">Ritesh's Notion</h1>
          </div>

          <div
            role="button"
            className={cn(
              "text-muted-foreground rounded-sm hover:bg-neutral-300 dark:hover:bg-neutral-600 absolute my-auto right-2 opacity-0 group-hover/sidebar:opacity-100 transition",
              isMobile && "opacity-100"
            )}
          >
            <ChevronsLeft className="h-7 w-7" />
          </div>
        </div>

        <div className="">
          <div className="flex items-center gap-2 py-2 px-2 hover:bg-gray-200 cursor-pointer">
            <Search className="text-neutral-600" />
            <span className="select-none">Search</span>
          </div>
          <div className="flex items-center gap-2 py-2 px-2 hover:bg-gray-200 cursor-pointer">
            <History className="text-neutral-600" />
            <span className="select-none">Updates</span>
          </div>
          <div className="flex items-center gap-2 py-2 px-2 hover:bg-gray-200 cursor-pointer">
            <Settings className="text-neutral-600" />
            <span className="select-none">Settings and members</span>
          </div>
          <div className="flex items-center gap-2 py-2 px-2 hover:bg-gray-200 cursor-pointer">
            <PlusCircle className="text-neutral-600" />
            <span className="select-none">New page</span>
          </div>
        </div>

        <div className="mt-4">
          <h1 className="font-semibold px-2 mb-2">Pages</h1>

          <div className="flex items-start gap-1 py-2 hover:bg-neutral-200 px-2 cursor-pointer">
            <ChevronRight className="text-neutral-400 h-6 w-6" />
            <span className="font-semibold select-none">🚞 My Document</span>
          </div>
        </div>

        {/* <div className="opacity-0 group-hover/sidebar:opacity-100 transition cursor-ew-resize absolute h-full w-1 bg-primary/10 right-0 top-0" /> */}
      </aside>

      <div
        ref={navbarRef}
        className={cn(
          "absolute top-0 z-[99999] left-60 w-[calc(100%-240px)]",
          isResetting && "transition-all ease-in-out duration-300",
          isMobile && "left-0 w-full"
        )}
      >
        <nav className="bg-transparent px-3 py-2 w-full">
          {isCollapsed && (
            <MenuIcon
              //   onClick={resetWidth}
              role="button"
              className="h-6 w-6 text-muted-foreground"
            />
          )}
        </nav>
      </div>
    </>
  )
}

export default Sidebar
