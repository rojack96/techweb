interface InsertMarkerProps {
    active: boolean
    onToggle: () => void
}

export function InsertMarker({ active, onToggle }: InsertMarkerProps) {
    return (
        <button
            onClick={onToggle}
            aria-label={active ? "Exit marker mode" : "Enter marker mode"}
            style={{
                position: "absolute",
                top: 12,
                right: 12,
                width: 44,
                height: 44,
                borderRadius: 8,
                border: "none",
                boxShadow: "0 3px 8px rgba(0,0,0,0.25)",
                backgroundColor: active ? "#2c3e50" : "#ff0000",
                color: "white",
                fontSize: 24,
                cursor: "pointer",
                zIndex: 1000,
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
            }}
        >
            {active ? (<svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="white"
                strokeWidth="3"
                strokeLinecap="round"
            >
                <line x1="6" y1="6" x2="18" y2="18" />
                <line x1="6" y1="18" x2="18" y2="6" />
            </svg>) : (
                <svg
                    width="20"
                    height="20"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="white"
                    strokeWidth="3"
                    strokeLinecap="round"
                >
                    <line x1="12" y1="5" x2="12" y2="19" />
                    <line x1="5" y1="12" x2="19" y2="12" />
                </svg>
            )}
        </button>
    )
}