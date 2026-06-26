import { apiClient } from "./client"

export const pairsApi = {
    getPairsNames: async () => {
        const res = await apiClient.get("/pairs/names")

        return res.data
    },
}