import { ReactNode } from "react"

const layout = ({ children }: { children: ReactNode }) => {
  return (
    <section className="flex bg-gray-100 justify-center flex-col items-center h-screen">
      <div className="mx-auto max-w-[450px] bg-white p-4 rounded-md">
        {children}
      </div>
    </section>
  )
}

export default layout
