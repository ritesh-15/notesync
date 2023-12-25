"use client"

import AuthService from "@/api/authService"
import { Button } from "@/components/ui/button"
import { toast } from "@/components/ui/use-toast"
import { Loader2, MailMinus } from "lucide-react"
import { useMutation } from "react-query"

interface IEmailSentProps {
  searchParams: {
    email: string
    name: string
  }
}

const EmailSent = ({ searchParams }: IEmailSentProps) => {
  const { email, name } = searchParams

  const loginMutation = useMutation(AuthService.login)

  const handleResendEmail = async () => {
    try {
      await loginMutation.mutateAsync({ email })
      toast({
        title: "Email resend successfully",
        description:
          "We have successfully sent you a new email for verification of account please check your inbox",
        duration: 3000,
      })
    } catch (e: any) {
      toast({
        title: "Uh oh! Something went wrong.",
        description:
          e.response?.data.message || "There was a problem with your request.",
        variant: "destructive",
        duration: 3000,
      })
    }
  }

  return (
    <div className="flex items-center justify-center flex-col">
      <MailMinus className="h-12 w-12" />
      <h1 className="text-xl font-bold mt-2">
        Vericiation email sent succesfully!
      </h1>
      <p className="text-center mt-2">
        Hi <span className="font-bold">{name}</span>, we have sent you a
        verification link on your <span className="font-bold">{email}</span>{" "}
        address please click on the link to procedd to your account
      </p>

      <Button
        disabled={loginMutation.isLoading}
        onClick={handleResendEmail}
        className="mt-4"
      >
        {loginMutation.isLoading ? (
          <Loader2 className="mr-2 h-4 w-4 animate-spin" />
        ) : (
          "  Didn't received email ? Resend"
        )}
      </Button>
    </div>
  )
}

export default EmailSent
