import { apiClient } from "./client"

export const spreadApi = {
    getSpread: async (pair) => {
        const res = await apiClient.get("/spread", {
            params: { pair },
        })

        return res.data
    },
}