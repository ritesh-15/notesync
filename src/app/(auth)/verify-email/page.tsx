import AuthService from "@/api/authService"
import { AxiosError } from "axios"
import { redirect } from "next/navigation"
import React from "react"

export interface IVerifyEmailData {
  token: string
  action: string
  userId: string
}

interface IVerifyEmailProps {
  searchParams: IVerifyEmailData
}

const verifyEmail = async (
  params: IVerifyEmailData
): Promise<string | null> => {
  try {
    const res = await AuthService.verifyEmail(params)
    return null
  } catch (e: any) {
    if (e instanceof AxiosError) {
      return e.response?.data.message
    } else {
      return "something weng wrong please try again later"
    }
  }
}

const VerifyEmail = async ({ searchParams }: IVerifyEmailProps) => {
  const res = await verifyEmail(searchParams)

  if (res == null) {
    redirect("/dashboard")
  }

  if (res) {
    return (
      <>
        <h1 className="font-bold text-center">Enable to verify your account</h1>
        <p className="text-center mx-auto mt-2">
          Possible error for not validating your account is
          {res}
        </p>
      </>
    )
  }

  return <div>Verifying your email address..</div>
}

export default VerifyEmail
