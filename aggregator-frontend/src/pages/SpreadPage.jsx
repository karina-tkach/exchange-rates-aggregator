import { useEffect, useState } from "react"
import { spreadApi } from "../api/spread"
import { pairsApi } from "../api/pairs"
import PairSelect from "../components/PairSelect"
import {useHandleError} from "../utils/useHandleError.js";
import LoadingPage from "./generic/LoadingPage.jsx";

export default function SpreadPage() {
    const [pairs, setPairs] = useState([])
    const [pair, setPair] = useState("")
    const [data, setData] = useState([])

    const [firstLoading, setFirstLoading] = useState(true)
    const [firstSpreadLoading, setFirstSpreadLoading] = useState(true)
    const [error, setError] = useState(null)

    const handleErr = useHandleError()

    const loadInitial = async () => {
        try {
            const res = await pairsApi.getPairsNames()

            const list = res.data ?? []
            setPairs(list)

            if (list.length > 0) {
                setPair(list[0])
            }
        } catch (err) {
            handleErr(err)
        } finally {
            setFirstLoading(false)
        }
    }

    useEffect(() => {
        loadInitial()
    }, [])

    const syncPairs = async () => {
        try {
            const res = await pairsApi.getPairsNames()
            const list = res.data ?? []

            setPairs(list)

            if (!list.includes(pair)) {
                setPair(list[0] ?? "")
            }

        } catch (err) {
            handleErr(err)
        }
    }

    const loadSpread = async (selectedPair) => {
        if (!selectedPair) return

        try {
            setError(null)

            const res = await spreadApi.getSpread(selectedPair)
            setData(res.data || [])
        } catch (err) {
            const status = err?.response?.status ?? 500
            const message = err?.response?.data?.message ?? err?.message ?? "Something went wrong"
            if (status === 422) {
                setData([])
                setError(message)
            } else {
                handleErr(err)
            }
        } finally {
            setFirstSpreadLoading(false)
        }
    }

    useEffect(() => {
        loadSpread(pair)
    }, [pair])

    useEffect(() => {
        if (!pair) return

        const id = setInterval(() => {
            syncPairs()
            loadSpread(pair)
        }, 30000)

        return () => clearInterval(id)
    }, [pair])

    if (firstLoading || firstSpreadLoading) return <LoadingPage/>
    return (
        <div className="min-h-screen bg-[#020617] text-[#DAFCE6] px-6 py-10">
            <div className="max-w-5xl mx-auto flex justify-between items-center">
                <h1 className="text-4xl font-bold text-[#14a3c7]">
                    Spread Terminal
                </h1>

                <PairSelect
                    value={pair}
                    onChange={setPair}
                    pairs={pairs}
                />
            </div>

            <div className="max-w-5xl mx-auto mt-8">
                {error && (
                    <div className="mb-4 bg-red-400 text-red-950 p-2 rounded-lg">
                        {error}
                    </div>
                )}

                {data && data.length !== 0 ? (
                    <div className="bg-[#00002B] border border-[#22267C] rounded-xl p-6">
                        <div className="text-[#14a3c7] text-2xl mb-4 font-bold underline">
                            Arbitrage Spread
                        </div>

                        <div className="grid grid-cols-2 gap-4 text-lg">
                            <div>
                                <span className="text-[#DAFCE6]/60">Buy Exchange:</span>
                                <div className="text-[#DAFCE6] font-medium">
                                    {data.buy_exchange}
                                </div>
                            </div>

                            <div>
                                <span className="text-[#DAFCE6]/60">Sell Exchange:</span>
                                <div className="text-[#DAFCE6] font-medium">
                                    {data.sell_exchange}
                                </div>
                            </div>

                            <div>
                                <span className="text-[#DAFCE6]/60">Buy Price:</span>
                                <div>{data.buy_price}</div>
                            </div>

                            <div>
                                <span className="text-[#DAFCE6]/60">Sell Price:</span>
                                <div>{data.sell_price}</div>
                            </div>
                        </div>

                        <div className="mt-6 text-3xl font-bold text-[#14a3c7]">
                            Spread: {data.spread_percent}%
                        </div>
                    </div>
                ) : (
                    <div className="bg-[#00002B] border border-[#22267C] rounded-xl p-6
                     text-center text-[#DAFCE6]/60 text-2xl">
                        No spread data available
                    </div>
                )}

                <div className="mt-3 text-lg text-[#DAFCE6]/40">
                    polling: 30s • pair: {pair}
                </div>
            </div>
        </div>
    )
}