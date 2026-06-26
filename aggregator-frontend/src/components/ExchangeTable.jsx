import { useState } from "react"

export default function ExchangeTable({exchanges, onSave}) {
    const [saving, setSaving] = useState(null)

    const handleToggle = async (exchange) => {
        setSaving(exchange.id)

        try {
            await onSave(exchange)
        } finally {
            setSaving(null)
        }
    }

    return (
        <div className="rounded-3xl overflow-hidden border border-[#22267C]">

            <table className="w-full">

                <thead className="bg-[#11134b]">
                <tr>
                    <th className="text-left px-6 py-4">Name</th>
                    <th className="text-center">Enabled</th>
                    <th className="text-center">Action</th>
                </tr>
                </thead>

                <tbody>
                {exchanges.map(exchange => (
                    <tr
                        key={exchange.id}
                        className="border-t border-[#22267C]"
                    >
                        <td className="px-6 py-4">
                            {exchange.name}
                        </td>

                        <td className="text-center">
                            {exchange.enabled
                                ? "🟢"
                                : "🔴"}
                        </td>

                        <td className="text-center">

                            <button
                                disabled={saving === exchange.id}
                                onClick={() => handleToggle(exchange)}
                                className={`px-4 py-2 rounded-xl transition hover:cursor-pointer
                                    ${
                                    exchange.enabled
                                        ? "bg-red-600 hover:bg-red-500"
                                        : "bg-green-600 hover:bg-green-500"
                                }`}
                            >
                                {saving === exchange.id
                                    ? "Saving..."
                                    : exchange.enabled
                                        ? "Disable"
                                        : "Enable"}
                            </button>
                        </td>
                    </tr>
                ))}
                </tbody>

            </table>

        </div>
    )
}