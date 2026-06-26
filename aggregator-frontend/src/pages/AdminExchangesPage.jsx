import { useEffect, useState } from "react"

import LoadingPage from "./generic/LoadingPage"

import Pagination from "../components/Pagination"
import ExchangeTable from "../components/ExchangeTable"

import { exchangesApi } from "../api/exchanges"
import { useHandleError } from "../utils/useHandleError"

export default function AdminExchangesPage() {
    const [loading, setLoading] = useState(true)

    const [page, setPage] = useState(1)

    const [data, setData] = useState({
        data: [],
        current_page: 1,
        last_page: 1,
    })

    const handleError = useHandleError()

    const load = async (selectedPage = page) => {
        try {
            const res = await exchangesApi.get(selectedPage)

            setData({
                data: res.data.data,
                current_page: res.data.meta.current_page,
                last_page: res.data.meta.last_page,
            })
        } catch (err) {
            handleError(err)
        } finally {
            setLoading(false)
        }

    }

    useEffect(() => {
        load(page)
    }, [page])

    const updateExchange = async (exchange) => {
        try {

            await exchangesApi.update(exchange.id, {
                is_enabled: !exchange.enabled,
            })

            setData(prev => ({
                ...prev,
                data: prev.data.map(item =>
                    item.id === exchange.id
                        ? {
                            ...item,
                            enabled: !item.enabled,
                        }
                        : item
                ),
            }))

        } catch (err) {
            handleError(err)
        }

    }

    if (loading)
        return <LoadingPage />

    return (
        <div className="min-h-screen bg-[#020617] text-[#DAFCE6]">
            <div className="max-w-7xl mx-auto px-8 py-10">
                <div className="mb-10">
                    <h1 className="text-4xl font-bold text-[#14a3c7]">
                        Exchange Management
                    </h1>

                    <p className="text-[#94a3b8] mt-2">
                        Enable or disable market data sources.
                    </p>

                </div>

                <ExchangeTable
                    exchanges={data.data}
                    onSave={updateExchange}
                />

                <Pagination
                    currentPage={data.current_page}
                    lastPage={data.last_page}
                    onChange={setPage}
                />
            </div>
        </div>
    )
}