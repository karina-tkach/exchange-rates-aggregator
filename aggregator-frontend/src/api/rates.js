import { apiClient } from "./client"

export const ratesApi = {
    getRates: async (pair) => {
        const res = await apiClient.get("/rates", {
            params: { pair },
        })

        return res.data
    },
}