import { ReactNode } from "react"

const layout = ({ children }: { children: ReactNode }) => {
  return (
    <section className="flex justify-center flex-col items-center h-screen">
      <div className="mx-auto w-full max-w-[400px] bg-white p-4 rounded-md">
        {children}
      </div>
    </section>
  )
}

export default layout
