"use client"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { signIn, useSession } from "next-auth/react"

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"
import Image from "next/image"
import { useToast } from "@/components/ui/use-toast"
import { Loader2 } from "lucide-react"
import { useState } from "react"
import { useRouter } from "next/navigation"

const loginSchema = z.object({
  email: z
    .string({ required_error: "email address is required" })
    .email("email address is not valid"),
})

type LoginSchema = z.infer<typeof loginSchema>

const Login = () => {
  const { toast } = useToast()
  const { data: user } = useSession({ required: false })
  const router = useRouter()

  const [isGoogleSignInLoading, setIsGoogleSignInLoading] = useState(false)

  if (user) router.replace("/")

  const form = useForm<LoginSchema>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
    },
  })

  const onSubmit = async (values: LoginSchema) => {
    try {
      await signIn("email", {
        email: values.email,
      })

      toast({
        title: "Email sent successfully",
        description: "Please check you inbox for the verification",
      })
    } catch (error: any) {
      toast({
        title: "Uh oh! Something went wrong.",
        description: "There was a problem with your request.",
        variant: "destructive",
      })
    }
  }

  const handleGoogleLogin = async () => {
    try {
      setIsGoogleSignInLoading(true)
      await signIn("google", { callbackUrl: "http://localhost:3000" })
    } catch (e: any) {
      toast({
        title: "Uh oh! Something went wrong.",
        description: "There was a problem with your request.",
        variant: "destructive",
      })
    } finally {
      setIsGoogleSignInLoading(false)
    }
  }

  return (
    <>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-center">Notesync</h1>
        <p className="text-center max-w-[75%] mx-auto">
          Get started by creating an account or loging in your account.
        </p>
      </div>

      <div className="mb-8">
        <Button
          onClick={handleGoogleLogin}
          variant="outline"
          className="w-full"
          disabled={isGoogleSignInLoading}
        >
          {isGoogleSignInLoading ? (
            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
          ) : (
            <>
              <Image
                src="/images/google_logo.png"
                className="mr-4"
                alt=""
                width={23}
                height={23}
              />
              Log in with Google
            </>
          )}
        </Button>
      </div>

      <hr className="mb-4" />

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email address</FormLabel>
                <FormControl>
                  <Input placeholder="johndoe@gmail.com" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <Button
            disabled={form.formState.isSubmitting}
            className="w-full"
            type="submit"
          >
            {form.formState.isSubmitting ? (
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            ) : (
              " Continue with email"
            )}
          </Button>
        </form>
      </Form>
    </>
  )
}

export default Login
