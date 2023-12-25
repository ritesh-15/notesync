import { RegisterSchema } from "../app/(auth)/register/page"
import { IVerifyEmailData } from "../app/(auth)/verify-email/page"
import api from "./axios"

class AuthService {
  static register(data: RegisterSchema) {
    return api.post("/auth/register", data)
  }

  static login(data: { email: string }) {
    return api.post("/auth/login", data)
  }

  static me() {
    return api.get("/auth/me")
  }

  static verifyEmail(data: IVerifyEmailData) {
    return api.post("/auth/verify", data)
  }
}

export default AuthService
