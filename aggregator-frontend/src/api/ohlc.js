import { apiClient } from "./client"

export const ohlcApi = {
    getCandles: async ({ pair, period, timeframe, exchange }) => {
        const params = { pair, period, timeframe }
        if (exchange && exchange !== "ALL") {
            params.exchange = exchange
        }

        const res = await apiClient.get("/ohlc", { params })

        return res.data
    },
}