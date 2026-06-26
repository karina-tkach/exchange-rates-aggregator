import { useEffect, useState } from "react"
import { ratesApi } from "../api/rates"
import { pairsApi } from "../api/pairs"
import PairSelect from "../components/PairSelect"
import RatesTable from "../components/RatesTable"
import {useHandleError} from "../utils/useHandleError.js";
import LoadingPage from "./generic/LoadingPage.jsx";

export default function RatesPage() {
    const [pairs, setPairs] = useState([])
    const [pair, setPair] = useState("")
    const [data, setData] = useState([])
    const [firstLoading, setFirstLoading] = useState(true)
    const [firstRateLoading, setFirstRateLoading] = useState(true)

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

    useEffect(() => {
        loadInitial()
    }, [])

    const loadRate = async (selectedPair) => {
        if (!selectedPair) return

        try {
            const res = await ratesApi.getRates(selectedPair)
            setData(res.data || [])
        } catch (err) {
            handleErr(err)
        } finally {
            setFirstRateLoading(false)
        }
    }

    useEffect(() => {
        loadRate(pair)
    }, [pair])

    useEffect(() => {
        if (!pair) return

        const id = setInterval(() => {
            syncPairs()
            loadRate(pair)
        }, 30000)

        return () => clearInterval(id)
    }, [pair])

    if (firstLoading || firstRateLoading) return <LoadingPage/>
    return (
        <div className="min-h-[83vh] bg-[#020617] text-[#DAFCE6] px-6 py-10">

            <div className="max-w-5xl mx-auto flex justify-between items-center">
                <h1 className="text-4xl font-bold text-[#14a3c7]">
                    Rates Terminal
                </h1>

                <PairSelect
                    value={pair}
                    onChange={setPair}
                    pairs={pairs}
                />
            </div>

            <div className="max-w-5xl mx-auto mt-8">
                {!data || data.length === 0 ? (
                    <div className="bg-[#00002B] border border-[#22267C] rounded-xl p-6
                     text-center text-[#DAFCE6]/60 text-2xl">
                        No market data available
                    </div>
                ) : (
                    <RatesTable data={data} />
                )}

                <div className="mt-3 text-lg text-[#DAFCE6]/40">
                    polling: 30s • pair: {pair}
                </div>

            </div>
        </div>
    )
}