import { useMapEvents } from "react-leaflet"

interface LocationMarkerProps {
    active: boolean
    onMarkerPlaced: (position: [number, number]) => void
}

export function LocationMarker({ active, onMarkerPlaced }: LocationMarkerProps) {
    useMapEvents({
        click(e) {
            if (!active) return
            const pos: [number, number] = [e.latlng.lat, e.latlng.lng]
            //console.log("Clicked at: ", pos)
            onMarkerPlaced(pos)
        },
    })

    return null
}