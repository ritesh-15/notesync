import { useForm, SubmitHandler } from "react-hook-form"

type LoginFormInput = {
  email: string
}

const useLogin = () => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<LoginFormInput>()

  const onSubmit: SubmitHandler<LoginFormInput> = (data) => console.log(data)

  return { handleSubmit, register, errors, onSubmit }
}

export default useLogin
