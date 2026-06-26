export default function Footer() {
    return (
        <footer className="bg-[#00002B] border-t border-[#14a3c7]">
            <div className="text-center py-8 text-gray-500 text-sm border-t border-[#22267C]">
                MarketFlow © {new Date().getFullYear()} — Built for high-frequency market data
            </div>
        </footer>
    );
}