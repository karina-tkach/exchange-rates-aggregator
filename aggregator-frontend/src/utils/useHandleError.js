import { useNavigate } from "react-router-dom"

export function useHandleError() {
    const navigate = useNavigate()

    return (err) => {
        const status = err?.response?.status ?? 500

        const message =
            err?.response?.data?.message ??
            err?.message ??
            "Something went wrong"

        navigate("/error", {
            state: {
                message,
                code: status,
            },
        })
    }
}