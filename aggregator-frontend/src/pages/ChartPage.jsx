import { useEffect, useMemo, useState, useRef } from "react"
import CandleChart from "../components/CandleChart"
import ChartToolbar from "../components/ChartToolbar"
import EmptyState from "../components/EmptyState"
import LoadingPage from "./generic/LoadingPage"

import { ohlcApi } from "../api/ohlc"
import { pairsApi } from "../api/pairs"
import { exchangesApi } from "../api/exchanges"
import {useHandleError} from "../utils/useHandleError.js";

export default function ChartPage() {
    const isSyncingRef = useRef(false)

    const [pairs, setPairs] = useState([])
    const [exchanges, setExchanges] = useState([])

    const [pair, setPair] = useState("")
    const [exchange, setExchange] = useState("")

    const [period, setPeriod] = useState("24h")
    const [timeframe, setTimeframe] = useState("1h")

    const [candles, setCandles] = useState([])

    const [isFirstLoad, setIsFirstLoad] = useState(true)
    const [chartLoading, setChartLoading] = useState(false)

    const handleErr = useHandleError()

    const timeframeOptions = useMemo(() => {
        switch (period) {
            case "1h":
                return [
                    { value: "1m", label: "1 Minute" },
                    { value: "5m", label: "5 Minutes" },
                ]

            case "24h":
                return [
                    { value: "1m", label: "1 Minute" },
                    { value: "5m", label: "5 Minutes" },
                    { value: "1h", label: "1 Hour" },
                    { value: "4h", label: "4 Hours" },
                ]

            case "7d":
                return [
                    { value: "1m", label: "1 Minute" },
                    { value: "5m", label: "5 Minutes" },
                    { value: "1h", label: "1 Hour" },
                    { value: "4h", label: "4 Hours" },
                    { value: "1d", label: "1 Day" },
                ]

            default:
                return []
        }
    }, [period])

    useEffect(() => {
        const valid = timeframeOptions.find(t => t.value === timeframe)

        if (!valid) {
            isSyncingRef.current = true
            setTimeframe(timeframeOptions[0]?.value ?? "1m")
        }
    }, [period])

    const syncInitial = async () => {
        try {
            const [pairsRes, exchangesRes] = await Promise.all([
                pairsApi.getPairsNames(),
                exchangesApi.getActiveExchanges(),
            ])

            const loadedPairs = pairsRes.data ?? []
            const loadedExchanges = exchangesRes.data ?? []

            setPairs(loadedPairs)
            setExchanges(loadedExchanges)

            if (!loadedPairs.includes(pair)) {
                setPair(loadedPairs[0] ?? "")
            }
        } catch (err) {
            handleErr(err)
        } finally {
            setIsFirstLoad(false)
        }
    }

    const loadChart = async () => {
        if (!pair) return

        if (isSyncingRef.current) {
            isSyncingRef.current = false
            return
        }

        try {
            const response = await ohlcApi.getCandles({
                pair,
                period,
                timeframe,
                exchange: exchange || undefined,
            })

            setCandles(response.data ?? [])
        } catch (err) {
            handleErr(err)
        }
    }

    useEffect(() => {
        const loadInitial = async () => {
            try {
                const [pairsRes, exchangesRes] = await Promise.all([
                    pairsApi.getPairsNames(),
                    exchangesApi.getActiveExchanges(),
                ])

                const loadedPairs = pairsRes.data ?? []
                const loadedExchanges = exchangesRes.data ?? []

                setPairs(loadedPairs)
                setExchanges(loadedExchanges)

                if (loadedPairs.length > 0) {
                    setPair(loadedPairs[0])
                }
            } catch (err) {
                handleErr(err)
            } finally {
                setIsFirstLoad(false)
            }
        }

        loadInitial()
    }, [])

    useEffect(() => {
        const load = async () => {
            if (!pair) return

            if (isSyncingRef.current) {
                isSyncingRef.current = false
                return
            }

            try {
                setChartLoading(true)
                const response = await ohlcApi.getCandles({
                    pair,
                    period,
                    timeframe,
                    exchange: exchange || undefined,
                })

                setCandles(response.data ?? [])
            } catch (err) {
                handleErr(err)
            } finally {
                setChartLoading(false)
            }
        }

        load()
    }, [pair, exchange, period, timeframe])

    useEffect(() => {
        if (!pair) return

        const id = setInterval(() => {
            syncInitial()
            loadChart()
        }, 30000)

        return () => clearInterval(id)
    }, [pair, exchange, period, timeframe])

    if (isFirstLoad) {
        return <LoadingPage />
    }

    return (
        <div className="min-h-screen bg-[#020617] text-[#DAFCE6]">
            <div className="mx-auto max-w-7xl px-8 py-10">

                <div className="mb-8">
                    <h1 className="text-4xl font-bold text-[#DAFCE6]">
                        Market Charts
                    </h1>

                    <p className="mt-2 text-[#94a3b8]">
                        Live OHLC candlestick charts with automatic updates every
                        30 seconds.
                    </p>
                </div>

                <ChartToolbar
                    pairs={pairs}
                    exchanges={exchanges}
                    pair={pair}
                    exchange={exchange}
                    period={period}
                    timeframe={timeframe}
                    timeframeOptions={timeframeOptions}
                    onPairChange={setPair}
                    onExchangeChange={setExchange}
                    onPeriodChange={setPeriod}
                    onTimeframeChange={setTimeframe}
                />

                <div className="mt-8 rounded-3xl border border-[#22267C] bg-[#00002B] p-6 shadow-xl">
                    <div className="mb-6 flex items-center justify-between">
                        <div>
                            <h2 className="text-2xl font-semibold">
                                {pair}
                            </h2>

                            <p className="mt-1 text-sm text-[#14a3c7]">
                                {exchange || "All Exchanges"}
                            </p>
                        </div>

                        <div className="rounded-xl border border-[#22267C] bg-[#020617] px-4 py-2 text-sm text-[#14a3c7]">
                            Refresh every 30 sec
                        </div>
                    </div>

                    {chartLoading ? (
                        <div className="flex h-[600px] items-center justify-center text-[#94a3b8]">
                            Loading chart...
                        </div>
                    ) : candles.length === 0 ? (
                        <div className="h-[600px]">
                            <EmptyState
                                title="No market data available"
                                description="No OHLC candles were found for the selected filters."
                            />
                        </div>
                    ) : (
                        <div className="h-[600px]">
                            <CandleChart data={candles} />
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}