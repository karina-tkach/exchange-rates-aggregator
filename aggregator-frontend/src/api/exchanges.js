import { apiClient } from "./client"

export const exchangesApi = {
    getActiveExchanges: async () => {
        const res = await apiClient.get("/exchanges/active")

        return res.data
    },

    get(page = 1) {
        return apiClient.get(`/exchanges?page=${page}`)
    },

    update(id, data) {
        return apiClient.patch(`/exchanges/${id}`, data)
    }
}