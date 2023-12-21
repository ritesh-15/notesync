import Sidebar from "@/components/dashboard/sidebar/Sidebar"
import getUser from "@/utils/getUser"
import { redirect } from "next/navigation"
import React, { ReactNode } from "react"

const MainLayout = async ({ children }: { children: ReactNode }) => {
  const session = await getUser()

  if (!session) redirect("/")

  return (
    <main className="h-screen flex w-full">
      <Sidebar />
      <section className="">{children}</section>
    </main>
  )
}

export default MainLayout
