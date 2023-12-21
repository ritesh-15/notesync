import Navbar from "@/components/navbar/Navbar"
import { Button } from "@/components/ui/button"
import Image from "next/image"
import Link from "next/link"

const Home = () => {
  return (
    <>
      <Navbar />
      <section className="flex items-center flex-col justify-center w-full h-full pt-16">
        <h1 className="text-5xl leading-normal font-bold max-w-[55%] sm:w-full text-center">
          Your Ideas, Documents and Plans. Unified. Welcome to NoteSync
        </h1>
        <p className="text-xl mt-4 max-w-[35%] mx-auto text-center sm:w-full">
          Notesync is the connected workspace where better, faster work happens.
        </p>

        <Link href="/dashboard">
          <Button size="lg" className="mt-6">
            Go to dashboard
          </Button>
        </Link>

        <div className="w-full max-w-[450px] h-[450px] relative">
          <Image src="/images/hero.jpg" alt="" fill />
        </div>
      </section>
    </>
  )
}

export default Home
