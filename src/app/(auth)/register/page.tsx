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
import { useMutation } from "react-query"
import AuthService from "@/api/authService"
import { AxiosError } from "axios"

const registerSchema = z.object({
  email: z
    .string({ required_error: "email address is required" })
    .email("email address is not valid"),
  name: z
    .string({ required_error: "name is required" })
    .trim()
    .min(3, "name must be greater than 3 characters"),
})

export type RegisterSchema = z.infer<typeof registerSchema>

const Register = () => {
  const { toast } = useToast()
  const router = useRouter()

  const searchParams = useSearchParams()
  const callbackURL = searchParams.get("callback")

  const form = useForm<RegisterSchema>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      email: "",
      name: "",
    },
  })

  const registerMutation = useMutation(AuthService.register)

  const onSubmit = async (values: RegisterSchema) => {
    try {
      await registerMutation.mutateAsync(values)

      router.push(`/email-sent?email=${values.email}&name=${values.name}`)

      toast({
        title: "Email sent successfully",
        description: "Please check you inbox for the verification",
        duration: 3000,
      })
    } catch (error: any) {
      toast({
        title: "Uh oh! Something went wrong.",
        description:
          error.response?.data.message ||
          "There was a problem with your request.",
        variant: "destructive",
        duration: 3000,
      })
    }
  }

  return (
    <>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-center">Notesync</h1>
        <p className="text-center mx-auto pt-1">
          Get started by creating an account with your email address
        </p>
      </div>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Name</FormLabel>
                <FormControl>
                  <Input placeholder="johndoe" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

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
          Already have an account ?{" "}
          <Link
            href={`/login?callback=${callbackURL}`}
            className="text-primary font-medium mr-2 text-md"
          >
            Login
          </Link>
        </p>
      </Form>
    </>
  )
}

export default Register
