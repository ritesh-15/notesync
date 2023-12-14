"use client"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import Image from "next/image"
import useLogin from "./useLogin"

const Login = () => {
  const { handleSubmit, register, errors, onSubmit } = useLogin()

  return (
    <section className="bg-gray-100 flex justify-center flex-col items-center h-screen">
      <div className="mx-auto max-w-[400px] bg-white p-4 rounded-md">
        <h1 className="text-3xl font-bold mb-6">Log In</h1>

        <Button variant="outline" className="w-full mb-8">
          <Image
            src="/images/google_logo.png"
            width={22}
            height={22}
            alt=""
            className="mr-2 aspect-square object-contain"
          />
          Login with Google
        </Button>

        <hr />

        <form onSubmit={handleSubmit(onSubmit)} className="w-full mt-6">
          <Label className="mb-2 block" htmlFor="email">
            Email
          </Label>
          <Input
            className="bg-gray-100"
            placeholder="johndoe@abc.com"
            {...register("email", { required: true })}
          />
          <Button type="submit" className="mt-4 w-full">
            Continue with email
          </Button>
        </form>

        <small className="mt-2 block text-center">
          By clicking you acknowledge that you have read and understood, and
          agree to Notesync
        </small>
      </div>
    </section>
  )
}

export default Login
