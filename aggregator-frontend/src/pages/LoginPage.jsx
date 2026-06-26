import { useForm } from "react-hook-form"
import { useAuth } from "../contexts/AuthContext.jsx"
import { useNavigate } from "react-router-dom"
import { useState } from "react"
import { login } from "../api/auth.js"

export function LoginPage() {
    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm()

    const [serverError, setServerError] = useState("")
    const [loading, setLoading] = useState(false)

    const navigate = useNavigate()
    const { setUser } = useAuth()

    const onSubmit = async (data) => {
        setServerError("")
        setLoading(true)

        try {
            const user = await login(data)
            setUser(user)
            navigate("/")
        } catch (err) {
            const status = err.response?.status || 500
            const message = err.response?.data?.message
                || err.response?.message || "Something went wrong"

            if (status === 422 || status === 401) {
                setServerError(message)
            } else {
                navigate("/error", {
                    state: { message, code: status }
                })
            }
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="min-h-[83vh] flex items-center justify-center bg-[#020617] text-[#DAFCE6] relative overflow-hidden px-6">
            <div className="absolute inset-0 bg-[radial-gradient(circle_at_center,rgba(20,163,199,0.08)_1px,transparent_1px)] [background-size:35px_35px]" />

            <div className="absolute w-[500px] h-[500px] bg-[#14a3c7]/10 blur-3xl rounded-full top-[-150px] left-[-150px]" />
            <div className="absolute w-[400px] h-[400px] bg-[#22267C]/20 blur-3xl rounded-full bottom-[-120px] right-[-120px]" />

            <form
                onSubmit={handleSubmit(onSubmit)}
                className="relative z-10 w-full max-w-md border border-[#22267C] bg-[#00002B]/90 backdrop-blur-xl rounded-xl shadow-2xl"
            >
                <div className="flex items-center justify-between px-4 py-3 border-b border-[#22267C] bg-[#020617]">
                    <div className="flex gap-2">
                        <span className="w-3 h-3 rounded-full bg-red-400" />
                        <span className="w-3 h-3 rounded-full bg-yellow-400" />
                        <span className="w-3 h-3 rounded-full bg-green-400" />
                    </div>

                    <div className="text-xs text-[#14a3c7] font-mono">
                        auth@terminal
                    </div>
                </div>

                <div className="p-6 font-mono">

                    <h2 className="text-xl mb-4 text-[#14a3c7]">
                        $ login --secure
                    </h2>

                    <p className="text-sm text-[#94a3b8] mb-6">
                        Enter credentials to access system
                    </p>

                    {serverError && (
                        <div className="mb-4 text-red-400 border border-red-500/40 bg-red-500/10 p-3 rounded">
                            error: {serverError}
                        </div>
                    )}

                    <div className="mb-4">
                        <label className="text-xs text-[#94a3b8]">
                            email
                        </label>

                        <input
                            type="email"
                            placeholder="user@terminal.io"
                            {...register("email", {
                                required: "email required"
                            })}
                            className="w-full mt-1 bg-transparent border border-[#22267C] rounded px-3 py-2 text-[#DAFCE6]
                            focus:border-[#14a3c7] outline-none"
                        />

                        {errors.email && (
                            <p className="text-xs text-red-400 mt-1">
                                {errors.email.message}
                            </p>
                        )}
                    </div>

                    <div className="mb-6">
                        <label className="text-xs text-[#94a3b8]">
                            password
                        </label>

                        <input
                            type="password"
                            placeholder="••••••••"
                            {...register("password", {
                                required: "password required"
                            })}
                            className="w-full mt-1 bg-transparent border border-[#22267C] rounded px-3 py-2 text-[#DAFCE6]
                            focus:border-[#14a3c7] outline-none"
                        />

                        {errors.password && (
                            <p className="text-xs text-red-400 mt-1">
                                {errors.password.message}
                            </p>
                        )}
                    </div>

                    <button
                        type="submit"
                        disabled={loading}
                        className="w-full py-2 rounded border border-[#14a3c7] text-[#14a3c7]
                        hover:bg-[#14a3c7]/10 hover:cursor-pointer transition disabled:opacity-50"
                    >
                        {loading ? "authenticating..." : "run login"}
                    </button>

                    <div className="mt-4 text-xs text-[#94a3b8]">
                        system: jwt-auth • mode: secure
                    </div>
                </div>
            </form>
        </div>
    )
}