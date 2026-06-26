import { apiClient } from "./client"

export const exchangesApi = {
    getActiveExchanges: async () => {
        const res = await apiClient.get("/exchanges/active")

        return res.data
    },
}