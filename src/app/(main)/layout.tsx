import Header from "@/components/dashboard/header/Header"
import Sidebar from "@/components/dashboard/sidebar/Sidebar"
import { redirect } from "next/navigation"
import React, { ReactNode } from "react"

const MainLayout = async ({ children }: { children: ReactNode }) => {
  return (
    <main className="h-screen flex w-full">
      <Sidebar />

      <section className="w-full">
        <Header />

        <div className="">{children}</div>
      </section>
    </main>
  )
}

export default MainLayout
