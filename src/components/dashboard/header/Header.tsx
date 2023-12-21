import {
  ChevronRight,
  ChevronLeft,
  PlusIcon,
  History,
  MessageSquare,
  Star,
  MoreHorizontal,
} from "lucide-react"

const Header = () => {
  return (
    <header className="flex items-center justify-between py-3 px-2  w-full">
      <div className="flex items-center gap-2">
        <ChevronLeft className="w-7 h-7 cursor-pointer hover:bg-neutral-100 rounded-sm" />
        <ChevronRight className="w-7 h-7 text-neutral-500 cursor-pointer hover:bg-neutral-100 rounded-sm" />
        <PlusIcon className="w-7 h-7 text-neutral-500 cursor-pointer hover:bg-neutral-100 rounded-sm" />

        <span>🚞 My Document</span>
      </div>

      <div className="flex items-center gap-3">
        <div className="p-1 rounded-sm hover:bg-neutral-100 cursor-pointer">
          <span className="text-md">Share</span>
        </div>

        <MessageSquare className="w-8 h-8 hover:bg-neutral-100 rounded-sm text-neutral-700 p-1 cursor-pointer" />
        <Star className="w-8 h-8 p-1 hover:bg-neutral-100 rounded-sm text-neutral-700 cursor-pointer" />
        <History className="w-8 h-8 p-1 hover:bg-neutral-100 rounded-sm text-neutral-700 cursor-pointer" />
        <MoreHorizontal className="w-8 h-8 p-1 hover:bg-neutral-100 rounded-sm text-neutral-700 cursor-pointer" />
      </div>
    </header>
  )
}

export default Header
