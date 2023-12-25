import { NextRequest, NextResponse } from "next/server"

export async function middleware(request: NextRequest) {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_SERVER_URL}/api/auth/me`,
      {
        method: "GET",
        credentials: "include",
      }
    ).then((res) => res.json())
    console.log(res)
  } catch (e) {
    console.log(e)
  }

  return NextResponse.next()
}

export const config = {
  matcher: ["/dashboard"],
}
