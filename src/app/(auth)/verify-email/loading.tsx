import { Loader2 } from "lucide-react"
import React from "react"

const loading = () => {
  return (
    <>
      <Loader2 className="mr-2 h-4 w-4 animate-spin" />
      <h1 className="text-lg mt-2">Verifying your account...</h1>
    </>
  )
}

export default loading
