"use client"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"

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
import { useToast } from "@/components/ui/use-toast"
import { Loader2 } from "lucide-react"
import { useRouter, useSearchParams } from "next/navigation"
import Link from "next/link"
import AuthService from "@/api/authService"
import { useMutation } from "react-query"

const loginSchema = z.object({
  email: z
    .string({ required_error: "email address is required" })
    .email("email address is not valid"),
})

type LoginSchema = z.infer<typeof loginSchema>

const Login = () => {
  const { toast } = useToast()
  const router = useRouter()

  const searchParams = useSearchParams()
  const callbackURL = searchParams.get("callback")

  const form = useForm<LoginSchema>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
    },
  })

  const loginMutation = useMutation(AuthService.login)

  const onSubmit = async (values: LoginSchema) => {
    try {
      const { data } = await loginMutation.mutateAsync(values)

      router.push(`/email-sent?email=${data.data.email}&name=${data.data.name}`)

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

  return (
    <>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-center">Notesync</h1>
        <p className="text-center mx-auto">
          Lets log you in to access you documents
        </p>
      </div>

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
            size="lg"
          >
            {form.formState.isSubmitting ? (
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            ) : (
              " Continue with email"
            )}
          </Button>
        </form>

        <p className="mt-4 text-sm text-center">
          Dont have an account yet ?{" "}
          <Link
            href={`/register?callback=${callbackURL}`}
            className="text-primary font-medium mr-2 text-md"
          >
            Register
          </Link>
        </p>
      </Form>
    </>
  )
}

export default Login
